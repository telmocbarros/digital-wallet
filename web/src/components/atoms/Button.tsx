export default function Button(props: {
  label: string;
  onClick: () => void;
  disabled?: boolean;
}) {
  const { onClick, label, disabled = false } = props;
  return (
    <button
      onClick={onClick}
      disabled={disabled}
      className={`px-4 py-2 rounded ${
        props.disabled
          ? 'bg-gray-300 text-gray-600 cursor-not-allowed'
          : 'bg-blue-500 text-white hover:bg-blue-600'
      }`}
    >
      {label}
    </button>
  );
}
