"use client";

import { Fragment, useEffect, useMemo, useState } from "react";
import Fuse, { IFuseOptions, RangeTuple } from "fuse.js";

import { useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { Word } from "@/types";

import { CheckBox } from "../CheckBox/CheckBox";
import { FuseSearchResult, getIndicesForKey, splitWordInTokens } from "./utils";

import "./Main.scss";

function RenderHighlights(props: { word: string; matches: RangeTuple[] }) {
  const { word, matches } = props;

  return (
    <>
      {splitWordInTokens(word, matches).map((t, i) => (
        <span className={t[1] ? "highlight" : ""} key={i}>
          {t}
        </span>
      ))}
    </>
  );
}

export function Main() {
  const [searchTerm, setSearchTerm] = useState("");

  const [searchResult, setSearchResult] = useState<FuseSearchResult[]>([]);

  const [filters, setFilters] = useState({
    byWord: true,
    byTranslation: false,
    byDescription: false,
  });

  // try get words from api

  const getWordsQuery = useQuery({
    queryKey: ["getWords"],
    queryFn: async () => {
      const res = await fetch("api/wordInfo");
      return (await res.json()) as Word[];
    },
  });

  const wordsData = getWordsQuery.data;

  const router = useRouter();

  const fuseSearch = useMemo(() => {
    if (!wordsData) {
      return undefined;
    }

    const filterKeys: string[] = [];

    if (filters.byWord) {
      filterKeys.push("word");
    }

    if (filters.byDescription) {
      filterKeys.push("description");
    }

    if (filters.byTranslation) {
      filterKeys.push("translations");
    }

    if (!filterKeys.length) {
      filterKeys.push("word");
    }

    const fuseOptions: IFuseOptions<unknown> = {
      includeScore: true,
      includeMatches: true,
      keys: filterKeys,
    };

    return new Fuse(wordsData, fuseOptions);
  }, [wordsData, filters]);

  useEffect(() => {
    if (!fuseSearch) {
      return;
    }

    setSearchResult(
      fuseSearch.search(searchTerm).map((r) => ({
        index: r.refIndex,
        item: r.item,
        matches: r.matches || [],
      }))
    );
  }, [searchTerm, fuseSearch]);

  const onWordClick = (word: string) => {
    // navigate to word
    router.push(`/wordInfo/${word}`);
  };

  const renderWordsList = () => {
    if (!wordsData) {
      return <></>;
    }

    if (!searchTerm.trim()) {
      return (
        <div className="wordList">
          {wordsData.map((w, i) => (
            <div onClick={() => onWordClick(w.word)} key={i}>
              {w.word}
            </div>
          ))}
        </div>
      );
    }

    // render words with highlights

    const renderHighlightsForTranslations = (r: FuseSearchResult) => {
      const translationsWithMatches: {
        word: string;
        matches: RangeTuple[];
      }[] = r.matches
        .filter((m) => m.key === "translations")
        .map((m) => ({
          word: m.value || "",
          matches: [...m.indices],
        }));

      return (
        <div className="translationHighlightContainer">
          {translationsWithMatches.map((m) => (
            <div key={m.word}>
              <RenderHighlights word={m.word} matches={m.matches} />
              <div className="overlay">{m.word}</div>
            </div>
          ))}
        </div>
      );
    };

    const renderDescriptionFragment = (r: FuseSearchResult) => {
      const descriptionFragment = r.matches.find(
        (m) => m.key === "description"
      );

      if (!descriptionFragment) {
        return null;
      }

      return (
        <div className="descriptionFragmentHighlight">
          <RenderHighlights
            word={descriptionFragment.value || ""}
            matches={[...descriptionFragment.indices]}
          />
          <div className="overlay">{descriptionFragment.value}</div>
        </div>
      );
    };

    return (
      <div className="wordList">
        {searchResult.map((r) => (
          <div key={r.index}>
            {filters.byWord && (
              <div className="wordHighlightContainer">
                <RenderHighlights
                  key={r.item.word}
                  word={r.item.word}
                  matches={getIndicesForKey("word", r.matches)}
                />
              </div>
            )}

            {filters.byTranslation && renderHighlightsForTranslations(r)}

            <div>
              <div onClick={() => onWordClick(r.item.word)}>{r.item.word}</div>
              {filters.byDescription && renderDescriptionFragment(r)}
            </div>
          </div>
        ))}
      </div>
    );
  };

  const toggleFilter = (key: keyof typeof filters) => {
    setFilters((p) => ({
      ...p,
      [key]: !filters[key],
    }));
  };

  return (
    <div className="MainPage">
      <h1>Dictionary search</h1>

      <div className="flex">
        <CheckBox
          label="word"
          checked={filters.byWord}
          onChange={() => toggleFilter("byWord")}
        />
        <CheckBox
          label="translation"
          checked={filters.byTranslation}
          onChange={() => toggleFilter("byTranslation")}
        />
        <CheckBox
          label="description"
          checked={filters.byDescription}
          onChange={() => toggleFilter("byDescription")}
        />
      </div>

      <input
        className="search"
        value={searchTerm}
        onChange={(ev) => {
          setSearchTerm(ev.target.value);
        }}
        placeholder="Search"
      />

      {renderWordsList()}
    </div>
  );
}
