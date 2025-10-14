import { useState } from 'react';
import Account from './Settings/Account';
import Appearance from './Settings/Appearance';
import Profile from './Settings/Profile';
import Security from './Settings/Security';

export default function SettingsPage() {
  const [activeTab, setActiveTab] = useState('Account');

  return (
    <div>
      <h1>Settings Page</h1>
      {/* Settings Header */}
      <ul className="flex flex-row space-between">
        <li className="mr-3.5">
          <a onClick={() => setActiveTab('Account')}>Account</a>
        </li>
        <li className="mr-3.5">
          <a onClick={() => setActiveTab('Profile')}>Profile</a>
        </li>
        <li className="mr-3.5">
          <a onClick={() => setActiveTab('Security')}>Security</a>
        </li>
        <li className="mr-3.5">
          <a onClick={() => setActiveTab('Appearance')}>Appearance</a>
        </li>
        {/* <li className="mr-3.5">
          <a onClick={() => setActiveTab('Notifications')}>Notifications</a>
        </li>
        <li className="mr-3.5">
          <a onClick={() => setActiveTab('Billing')}>Billing</a>
        </li>
        <li className="mr-3.5">
          <a onClick={() => setActiveTab('Integration')}>Integration</a>
        </li> */}
      </ul>

      {/* Chosen Settings View */}
      <div className="max-h-[600px] overflow-y-auto">
        <Account className={activeTab === 'Account' ? '' : 'hidden'} />
        <Profile className={activeTab === 'Profile' ? '' : 'hidden'} />
        <Security className={activeTab === 'Security' ? '' : 'hidden'} />
        <Appearance className={activeTab === 'Appearance' ? '' : 'hidden'} />
      </div>
    </div>
  );
}
