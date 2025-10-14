import Account from './Settings/Account';
import Appearance from './Settings/Appearance';
import Profile from './Settings/Profile';

export default function SettingsPage() {
  return (
    <div>
      <h1>Settings Page</h1>
      {/* Settings Header */}
      <ul className="flex flex-row space-between">
        <li className="mr-3.5">
          <a>Account</a>
        </li>
        <li className="mr-3.5">
          <a>Profile</a>
        </li>
        <li className="mr-3.5">
          <a>Security</a>
        </li>
        <li className="mr-3.5">
          <a>Appearance</a>
        </li>
        <li className="mr-3.5">
          <a>Notifications</a>
        </li>
        <li className="mr-3.5">
          <a>Billing</a>
        </li>
        <li className="mr-3.5">
          <a>Integration</a>
        </li>
      </ul>

      {/* Chosen Settings View */}
      <div className="max-h-[600px] overflow-y-auto">
        <Account />
        {/* <Profile /> */}
        {/* <Appearance /> */}
      </div>
    </div>
  );
}
