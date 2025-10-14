import SearchBar from './SearchBar';

export default function TableHeader(props: {
  columns: Array<{ label: string; sortable?: boolean; onSort?: () => void }>;
}) {
  return (
    <div className="flex flex-row items-center border-b border-gray-300 pb-2 mb-2" role="row">
      {props.columns &&
        props.columns.map((column, index) => (
          <div key={index} className="text-left flex-2 font-semibold" role="columnheader">
            {column.sortable ? (
              <button
                type="button"
                className="hover:text-indigo-500 cursor-pointer transition-colors"
                onClick={column.onSort}
              >
                {column.label}
              </button>
            ) : (
              <span className="hover:text-indigo-500 cursor-pointer transition-colors">
                {column.label}
              </span>
            )}
          </div>
        ))}
      <SearchBar className="flex-1" />
    </div>
  );
}
