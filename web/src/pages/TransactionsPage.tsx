import TransactionsTable from '../components/organisms/transactions/TransactionsTable';

export default function Transactions() {
  const tableColumns = [
    { label: 'Date', sortable: true, onSort: () => {} },
    { label: 'Description', sortable: true, onSort: () => {} },
    { label: 'Amount', sortable: true, onSort: () => {} },
    { label: 'Category', sortable: true, onSort: () => {} },
  ];

  const transactions = [
    { id: '1', amount: 50, date: '2024-06-01' },
    { id: '2', amount: 75, date: '2024-06-02' },
    { id: '3', amount: 100, date: '2024-06-03' },
    { id: '4', amount: 200, date: '2024-06-04' },
    { id: '5', amount: 150, date: '2024-06-05' },
    { id: '6', amount: 300, date: '2024-06-06' },
    { id: '7', amount: 250, date: '2024-06-07' },
    { id: '8', amount: 400, date: '2024-06-08' },
    { id: '9', amount: 350, date: '2024-06-09' },
    { id: '10', amount: 500, date: '2024-06-10' },
  ];

  const pagination = 'Page 1 of 10';
  return (
    <div>
      <h1>Transactions Page</h1>
      <TransactionsTable
        columns={tableColumns}
        transactions={transactions}
        pagination={pagination}
      />
      {/**
       * list of transactions: TRANSACTIONS TABLE
       * filter by date, amount, category - TRANSACTIONS HEADER | FILTER
       * search by description - TRANSACTIONS HEADER | SEARCH
       * pagination - TRANSACTIONS FOOTER | PAGINATION
       */}
    </div>
  );
}
