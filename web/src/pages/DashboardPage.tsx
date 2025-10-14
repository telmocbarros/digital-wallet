import BalanceCard from '../components/molecules/BalanceCard';
import TableHeader from '../components/molecules/TableHeader';
import CardsList from '../components/organisms/CardsList';
import TransactionsList from '../components/organisms/transactions/TransactionsList';

export default function DashboardPage() {
  return (
    <div className="flex flex-row h-full py-2.5">
      <div className="flex-2 flex flex-col">
        <h1 className="text-2xl font-bold mb-4">My Cards</h1>
        <CardsList className="mb-4" />
        <BalanceCard className="mb-4" />
        <button className="mx-auto mt-4 p-2 bg-blue-500 text-white rounded">
          Add New Card
        </button>
      </div>
      <div className="flex-3">
        <h1 className="text-2xl font-bold">My Payments</h1>
        <TableHeader
          columns={[{ label: 'Recent Transactions' }, { label: 'View All' }]}
        />
        <TransactionsList />
        <h1 className="text-2xl font-bold mt-8">Update Payments</h1>
        <TransactionsList />
      </div>
    </div>
  );
}
