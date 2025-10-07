import { type ReactNode, useState } from 'react';
import './MainLayout.css';
import Navbar from '../organisms/Navbar';
import SlideMenu from '../organisms/SlideMenu';

interface MainLayoutProps {
  children: ReactNode;
}

export default function MainLayout({ children }: MainLayoutProps) {
  const [displayAppMenu, setDisplayAppMenu] = useState(false);
  const [displayProfileMenu, setDisplayProfileMenu] = useState(false);

  function handleAppIconClick() {
    setDisplayAppMenu(!displayAppMenu);
    setDisplayProfileMenu(false); // Close other menu
  }

  function handleProfileIconClick() {
    setDisplayProfileMenu(!displayProfileMenu);
    setDisplayAppMenu(false); // Close other menu
  }

  return (
    <>
      <Navbar
        onAppIconClick={handleAppIconClick}
        onProfileIconClick={handleProfileIconClick}
      />

      <SlideMenu isOpen={displayAppMenu} side="left">
        <ul>
          <li className="font-extrabold">Dashboard</li>
          <li>Transactions</li>
          <li>Wallet</li>
        </ul>
      </SlideMenu>

      <SlideMenu isOpen={displayProfileMenu} side="right">
        <ul>
          <li>Profile</li>
          <li>Settings</li>
          <li>Logout</li>
        </ul>
      </SlideMenu>

      <div className="content">{children}</div>
    </>
  );
}
