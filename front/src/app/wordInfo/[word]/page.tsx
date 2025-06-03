'use client';

import { useEffect, useMemo, useState } from 'react';

import { Word } from '@/types';
import { useQuery } from '@tanstack/react-query';
import Link from 'next/link';

import './page.scss';

export default function WordInfo({
  params,
}: {
  params: Promise<{
    word: string;
  }>;
}) {
  const [word, setWord] = useState<string>();

  useEffect(() => {
    params.then((w) => setWord(decodeURIComponent(w.word)));
  }, [params]);

  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);

  const [deleteConfirmed, setDeleteConfirmed] = useState(false);

  const getWordQuery = useQuery({
    queryKey: ['getWordInfo', word],
    queryFn: async () => {
      const res = await fetch(`/api/wordInfo/${word}`, {
        headers: {
          InitData: Telegram.WebApp.initData,
        },
      });
      return (await res.json()) as Word;
    },
    enabled: !!word,
  });

  const deleteWordQuery = useQuery({
    queryKey: ['deleteWord', word],
    queryFn: async () => {
      const res = await fetch(`/api/wordInfo/${word}`, {
        method: 'DELETE',
        headers: {
          InitData: Telegram.WebApp.initData,
        },
      });

      return res.status;
    },
    enabled: deleteConfirmed,
  });

  const showActions = useMemo(() => {
    if (showDeleteConfirm) {
      return false;
    }

    if (deleteConfirmed && deleteWordQuery.data == 200) {
      return false;
    }

    return true;
  }, [deleteConfirmed, showDeleteConfirm, deleteWordQuery.data]);

  useEffect(() => {
    if (deleteWordQuery.data === 200) {
    }
  }, [deleteWordQuery.data]);

  if (getWordQuery.isLoading) {
    return (
      <div className="WordInfoPage">
        <div>Loading</div>
      </div>
    );
  }

  const wordData = getWordQuery.data as Word;

  const deleteWord = () => {
    setShowDeleteConfirm(false);
    setDeleteConfirmed(true);
  };

  if (!wordData) {
    return <div>Loading</div>;
  }

  return (
    <div className="WordInfoPage">
      <div className="content">
        <h1>{word}</h1>
        <div className="translations">
          {wordData.translations.map((t, i) => (
            <div key={i}>{t}</div>
          ))}
        </div>
        <p>{wordData.description}</p>
        {!!wordData.similar?.length && (
          <>
            <h4>Group:</h4>
            {wordData.similar.map((w, i) => (
              <h3 key={i}>{w}</h3>
            ))}
          </>
        )}
      </div>

      <div className="footer">
        {showActions && (
          <div className="buttons">
            <Link href={`/editWord/${word}`}>Edit</Link>
            <button
              onClick={() => setShowDeleteConfirm(true)}
              className="destructive"
            >
              Delete
            </button>
          </div>
        )}

        {showDeleteConfirm && (
          <>
            <div>Are you sure?</div>
            <div className="buttons">
              <button onClick={() => setShowDeleteConfirm(false)}>
                Cancel
              </button>
              <button onClick={deleteWord} className="destructive">
                Delete!
              </button>
            </div>
          </>
        )}

        {deleteWordQuery.data === 200 && <h3>The word {word} was deleted!</h3>}
      </div>
    </div>
  );
}
