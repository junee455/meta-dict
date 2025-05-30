'use client';

import { useEffect, useState } from 'react';

import { EditWord } from '@/shared/components';
import { Word } from '@/types';
import { useQuery } from '@tanstack/react-query';

export default function EditWordPage({
  params,
}: {
  params: Promise<{ word: string }>;
}) {
  const [word, setWord] = useState<string>();

  useEffect(() => {
    params.then((w) => setWord(w.word));
  }, [params]);

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

  if (getWordQuery.isLoading || !getWordQuery.data) {
    return <div>Loading...</div>;
  }

  return <EditWord initialWord={getWordQuery.data} update />;
}
