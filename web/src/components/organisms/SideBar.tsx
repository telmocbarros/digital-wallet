import NavItem from '../molecules/NavItem';

export default function SideBar() {
  return (
    <div className="h-full flex flex-col justify-between">
      <div>
        <NavItem name="home" label="Home" />
        <NavItem name="transactions" label="Transactions" />
        <NavItem name="invoice" label="Invoices" />
        <NavItem name="settings" label="Settings" />
      </div>
      <div>
        <NavItem name="help" label="Help" />
        <NavItem name="settings" label="Logout" />
      </div>
    </div>
  );
}
