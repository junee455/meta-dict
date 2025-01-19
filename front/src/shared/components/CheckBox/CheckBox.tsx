import "./CheckBox.scss";

export type CheckBoxProps = {
  label?: string;
  checked?: boolean;
  onChange?: () => void;
};

export function CheckBox({ label, checked, onChange }: CheckBoxProps) {
  return (
    <div className="CheckBox">
      <button onClick={onChange} className="checkboxButton">
        {checked ? "✔" : ""}
      </button>
      <div>{label}</div>
    </div>
  );
}
