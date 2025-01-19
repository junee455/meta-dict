export type Word = {
  id: string;
  word: string;
  description: string;
  translations: string[];
  metadata: unknown;
  similar: string[];
};
