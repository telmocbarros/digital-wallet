export default function SettingsPage() {
  return (
    <div>
      <h1>Settings Page</h1>
      {/* Settings Header */}
      <ul className="flex flex-row space-between">
        <li className="mr-3.5">
          <a className="appearance-none">Account</a>
        </li>
        <li className="mr-3.5">
          <a className="appearance-none">Profile</a>
        </li>
        <li className="mr-3.5">
          <a className="appearance-none">Security</a>
        </li>
        <li className="mr-3.5">
          <a className="appearance-none">Appearance</a>
        </li>
        <li className="mr-3.5">
          <a className="appearance-none">Notifications</a>
        </li>
        <li className="mr-3.5">
          <a className="appearance-none">Billing</a>
        </li>
        <li className="mr-3.5">
          <a className="appearance-none">Integration</a>
        </li>
      </ul>

      {/* Chosen Settings View */}
      <div></div>
    </div>
  );
}
