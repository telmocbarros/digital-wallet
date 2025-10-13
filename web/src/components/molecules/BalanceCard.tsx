import Text from '../atoms/Text';
export default function BalanceCard({ className }: { className?: string }) {
  return (
    <div className={`bg-white shadow-md rounded-lg p-4 ${className}`}>
      <div className="">
        <Text className="text-gray-500">Total Balance</Text>
        <div className="flex">
          <Text className="flex-2 text-2xl font-bold">$5,430.00</Text>
          <Text className="text-sm text-gray-500">⬆ 23.04%</Text>
          <Text className="text-sm text-gray-500">⬇ 10.40%</Text>
        </div>
        <div className="flex">
          <div className="flex-2">
            <Text className="text-sm text-gray-500">Currency</Text>
            <Text className="flex-2 text-2xl font-bold">€ | EUR</Text>
          </div>
          <div>
            <Text className="text-sm text-gray-500">Status</Text>
            <Text className="flex-2 text-2xl font-bold">Active</Text>
          </div>
        </div>
      </div>
      <div></div>
    </div>
  );
}
