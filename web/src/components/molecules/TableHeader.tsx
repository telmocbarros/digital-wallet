// import SearchBar from './SearchBar';

export default function TableHeader(props: { children: React.ReactNode[] }) {
  return (
    <div className="flex flex-row items-center border-b-1 border-gray-300">
      <div className="flex flex-row flex-2">
        {props.children.map((child) => (
          <div className="mr-2.5 hover:border-b-3 hover:border-indigo-500 cursor-pointer">
            {child}
          </div>
        ))}
      </div>
      {/* <SearchBar className="flex-1" /> */}
    </div>
  );
}
