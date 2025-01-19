import { ChangeEvent, useEffect, useState } from "react";

export function useInput(initialValue = "") {
  const [value, setValue] = useState(initialValue);

  useEffect(() => {
    if (initialValue) {
      setValue(initialValue);
    }
  }, [initialValue]);

  const onChange = (
    event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    setValue(event.target.value);
  };

  return {
    value,
    onChange,
  };
}
