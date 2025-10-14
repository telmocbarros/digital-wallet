import SearchBar from './SearchBar';
export default function TableHeader(props: {
  columns: Array<{ label: string; sortable?: boolean; onSort?: () => void }>;
}) {
  return (
    <table className="w-full">
      <thead>
        <tr className="flex flex-row border-b border-gray-300">
          {props.columns &&
            props.columns.map((column, index) => (
              <th key={index} className="text-left flex-2">
                {column.sortable ? (
                  <button
                    type="button"
                    className="hover:text-indigo-500 cursor-pointer"
                    onClick={column.onSort}
                  >
                    {column.label}
                  </button>
                ) : (
                  <span className="hover:border-b-1 hover:border-indigo-500 cursor-pointer">
                    {column.label}
                  </span>
                )}
              </th>
            ))}
          <SearchBar className="flex-1" />
        </tr>
      </thead>
    </table>
  );
}
