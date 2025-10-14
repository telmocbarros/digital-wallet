export function TableFooter(props: { children: React.ReactNode }) {
  return (
    <div className="flex flex-row justify-end items-center border-t border-gray-300 mt-4 pt-4">
      {props.children}
    </div>
  );
}
