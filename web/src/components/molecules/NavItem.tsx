import Icon from '../atoms/icons/Icon';
import Text from '../atoms/Text';
export default function NavItem({
  name,
  label,
}: {
  name: string;
  label: string;
}) {
  return (
    <div className="flex flex-row items-center gap-2 p-2 hover:bg-gray-700 rounded-lg cursor-pointer">
      <Icon className="w-6 h-6" name={name} />
      <Text>{label}</Text>
    </div>
  );
}
