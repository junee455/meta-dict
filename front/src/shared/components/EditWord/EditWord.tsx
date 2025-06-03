'use client';

import { useState } from 'react';

import { useInput } from '@/shared/hooks';
import { Word } from '@/types';
import { useQuery } from '@tanstack/react-query';
import { useRouter } from 'next/navigation';

import './EditWord.scss';

export type EditWordProps = {
  initialWord?: Word;
  update?: boolean;
};

export function EditWord(props: EditWordProps) {
  const { initialWord, update = false } = props;

  const wordInputHook = useInput(initialWord?.word);
  const descriptionInputHook = useInput(initialWord?.description);
  const [shouldFetch, setShouldFetch] = useState(false);

  const [editingTranslation, setEditingTranslation] = useState<{
    value: string;
    oldValue: string;
    existingIndex?: number;
  }>();

  const [translations, setTranslations] = useState(
    initialWord?.translations || []
  );

  const router = useRouter();

  useQuery({
    queryKey: ['addWord'],
    queryFn: async () => {
      let reqMethod: string;

      if (update) {
        reqMethod = 'PATCH';
      } else {
        reqMethod = 'POST';
      }

      const wordToSave = update ? initialWord?.word : wordInputHook.value;

      const res = await fetch('/api/wordInfo', {
        method: reqMethod,
        body: JSON.stringify({
          translations,
          word: wordToSave,
          description: descriptionInputHook.value,
          similar: [],
        }),
        headers: {
          InitData: Telegram.WebApp.initData,
        },
      }).then((r) => {
        router.replace(`/wordInfo/${wordToSave}`);
        return r;
      });

      return res.status;
    },
    enabled: shouldFetch,
  });

  const saveWord = () => {
    // validate data
    const readyToFetch =
      !!wordInputHook.value.trim() &&
      (!!descriptionInputHook.value.trim() || !!translations.length);

    setShouldFetch(readyToFetch);
  };

  const editExistingTranslation = (index: number) => {
    setEditingTranslation({
      value: translations[index],
      oldValue: translations[index],
      existingIndex: index,
    });
  };

  const editTranslationInput = useInput(editingTranslation?.oldValue);

  const cancelEditTranslation = () => {
    setEditingTranslation(undefined);
  };

  const saveTranslation = () => {
    if (editingTranslation?.existingIndex !== undefined) {
      // remove old translation if empty
      if (!editTranslationInput.value.trim()) {
        setTranslations((p) =>
          p.filter((_, i) => i !== editingTranslation.existingIndex)
        );
      } else {
        // update existing one
        setTranslations((p) => {
          const newTranslations = [...p];
          newTranslations[editingTranslation.existingIndex!] =
            editTranslationInput.value;

          return newTranslations;
        });
      }
    } else {
      // add new translation
      if (!!editTranslationInput.value.trim()) {
        setTranslations((p) => [...p, editTranslationInput.value]);
      }
    }

    setEditingTranslation(undefined);
  };

  const addNewTranslation = () => {
    editTranslationInput.setValue('');
    setEditingTranslation({ value: '', oldValue: '' });
  };

  return (
    <div className="EditWord">
      <div className="content">
        {update && <h1>Edit word</h1>}
        {!update && <h1>Add new word</h1>}

        <h4>Word</h4>
        <input {...wordInputHook.inputProps} />

        <h4>Translations</h4>
        <div className="translations">
          {translations.map((t, i) => (
            <div onClick={() => editExistingTranslation(i)} key={t}>
              {t}
            </div>
          ))}
        </div>

        {!editingTranslation && (
          <button onClick={addNewTranslation} className="addTranslation">
            Add another translation
          </button>
        )}

        {!!editingTranslation && (
          <div className="flex">
            <input {...editTranslationInput.inputProps} />
            <button onClick={saveTranslation}>✔</button>
            <button onClick={cancelEditTranslation}>✖</button>
          </div>
        )}

        <h4>Description</h4>
        <textarea {...descriptionInputHook.inputProps} rows={10} />
      </div>

      <div className="footer">
        <button onClick={saveWord}>Save</button>
      </div>
    </div>
  );
}
