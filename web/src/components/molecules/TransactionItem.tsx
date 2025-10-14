import Icon from '../atoms/icons/Icon';
import Text from '../atoms/Text';

export default function TransactionItem() {
  return (
    <div
      className="flex flex-row justify-between items-center py-4 border-b border-gray-100 hover:bg-gray-50 transition-colors"
      role="row"
    >
      <div className="flex flex-row items-center gap-3 flex-2" role="cell">
        <Icon className="w-10 h-10" />
        <img src="../../../assets/facebook.png" alt="Transaction logo" className="w-10 h-10 rounded" />
        <div className="flex flex-col">
          <Text className="font-medium">Payoneer</Text>
          <Text className="text-sm text-gray-500">22 Apr 2022, 09:00 AM</Text>
        </div>
      </div>
      <div className="font-semibold text-green-600" role="cell">+ $4500.00</div>
    </div>
  );
}
