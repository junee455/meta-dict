import { Word } from "@/types";
import { FuseResultMatch, RangeTuple } from "fuse.js";

export type FuseSearchResult = {
  index: number;
  item: Word;
  matches: readonly FuseResultMatch[];
};

export function getIndicesForKey(
  key: string,
  matches: readonly FuseResultMatch[]
): RangeTuple[] {
  return matches
    .filter((m) => m.key === key)
    .map((m) => [...m.indices])
    .reduce((prev, curr) => [...prev, ...curr], [] as RangeTuple[]);
}

export function splitWordInTokens(word: string, matches: number[][]) {
  const matchRanges = matches.map((m) => [m[0], m[1] + 1]);

  let wordLeftover = word;

  const tokens: [string, boolean][] = [];

  matchRanges.forEach((mRange) => {
    const highlightedToken = word.slice(...mRange);

    const splitRes = wordLeftover.split(highlightedToken);

    const dimToken = splitRes[0];
    wordLeftover = splitRes.slice(1).join(highlightedToken);

    if (wordLeftover === undefined) {
      wordLeftover = "";
    }

    tokens.push([dimToken, false], [highlightedToken, true]);
  });

  tokens.push([wordLeftover, false]);

  return tokens.filter((t) => !!t[0]);
}
