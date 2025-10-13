import Icon from '../atoms/icons/Icon';
import Text from '../atoms/Text';

export default function TransactionItem() {
  return (
    <div className="flex flex-row justify-between items-center py-4 ">
      <Icon className="w-13 h-13" />
      <img src="../../../assets/facebook.png" alt="" />
      <div className="flex-2">
        <Text>Payoneer</Text>
        <Text>22 Apr 2022, 09:00 AM</Text>
      </div>
      <div>+ $4500.00</div>
    </div>
  );
}
