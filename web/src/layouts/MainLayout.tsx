import { type ReactNode } from 'react';
import './MainLayout.css';
import DigitalWalletIcon from '../shared/components/icons/DigitalWalletIcon';
import UserIcon from '../shared/components/icons/UserIcon';
import { useState } from 'react';

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
      <nav className="main-navbar">
        <div className="navbar-brand">
          <DigitalWalletIcon
            onClick={handleAppIconClick}
            className="navbar-icon"
          />
          <span className="navbar-title">Cozy Wallet</span>
        </div>
        <div className="navbar-actions">
          <UserIcon onClick={handleProfileIconClick} className="navbar-icon" />
        </div>
      </nav>
      {/* Left menu - App navigation */}
      <div className={`slide-menu slide-menu--left ${displayAppMenu ? 'open' : ''}`}>
        <ul>
          <li>Dashboard</li>
          <li>Transactions</li>
          <li>Wallet</li>
          <li>Settings</li>
        </ul>
      </div>

      {/* Right menu - Profile */}
      <div className={`slide-menu slide-menu--right ${displayProfileMenu ? 'open' : ''}`}>
        <ul>
          <li>Profile</li>
          <li>Settings</li>
          <li>Logout</li>
        </ul>
      </div>

      <div className="content">{children}</div>
    </>
  );
}
