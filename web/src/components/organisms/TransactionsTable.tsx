import { TableFooter } from '../molecules/TableFooter';
import TableHeader from '../molecules/TableHeader';
import TransactionItem from '../molecules/TransactionItem';

export default function TransactionsTable(props: {
  filters: { label: string; sortable?: boolean; onSort?: () => void }[];
  transactions: { id: string; amount: number; date: string }[];
  pagination: string | number;
}) {
  return (
    <div>
      <TableHeader columns={props.filters} />
      <div>
        <TransactionItem />
      </div>
      <TableFooter>{props.pagination}</TableFooter>
    </div>
  );
}
