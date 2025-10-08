import type { JSX } from 'react';
import UserIcon from '../atoms/icons/UserIcon';

function UserBarIcon(
  props: JSX.IntrinsicAttributes & React.HTMLAttributes<HTMLDivElement>
) {
  const { className, ...rest } = props;
  return (
    <div
      {...rest}
      className={`flex items-center gap-2 cursor-pointer ${className}`}
    >
      <UserIcon
        onClick={() => console.log('User icon clicked')}
        className="navbar-icon"
      />
      <div className="flex flex-col">
        <span className="text-xs font-semibold">User Name</span>
        <span className="text-xs">User Email</span>
      </div>
      <div>
        <span>▼</span>
      </div>
    </div>
  );
}
export default UserBarIcon;
