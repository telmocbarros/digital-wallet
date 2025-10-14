import { TableFooter } from '../../molecules/TableFooter';
import TableHeader from '../../molecules/TableHeader';
import TransactionItem from '../../molecules/TransactionItem';

export default function TransactionsTable(props: {
  columns: { label: string; sortable?: boolean; onSort?: () => void }[];
  transactions: { id: string; amount: number; date: string }[];
  pagination: string | number;
  maxHeight?: string; // e.g., "400px", "50vh", "max-h-96"
}) {
  const heightClass = props.maxHeight?.startsWith('max-h-')
    ? props.maxHeight
    : '';

  const inlineStyle =
    props.maxHeight && !props.maxHeight.startsWith('max-h-')
      ? { maxHeight: props.maxHeight }
      : {};

  return (
    <div
      className="w-full flex flex-col"
      role="table"
      aria-label="Transactions"
    >
      <TableHeader columns={props.columns} />
      <div
        className={`overflow-y-auto ${heightClass || 'max-h-[500px]'}`}
        style={inlineStyle}
        role="rowgroup"
      >
        {props.transactions.map((transaction) => (
          <TransactionItem key={transaction.id} />
        ))}
      </div>
      <TableFooter>{props.pagination}</TableFooter>
    </div>
  );
}
