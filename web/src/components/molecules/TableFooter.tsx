export function TableFooter(props: { children: React.ReactNode }) {
  return (
    <div>
      <span>&lt;</span>
      <div className="flex flex-row justify-end border-t-1 border-gray-300 mt-4 pt-4">
        {props.children}
      </div>
      <span>&gt;</span>
    </div>
  );
}
