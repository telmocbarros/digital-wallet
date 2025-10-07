import { type ReactNode } from 'react';
import './MainLayout.css';
import DigitalWalletIcon from '../shared/components/icons/DigitalWalletIcon';
import UserIcon from '../shared/components/icons/UserIcon';

interface MainLayoutProps {
  children: ReactNode;
}

export default function MainLayout({ children }: MainLayoutProps) {
  return (
    <>
      <nav className="main-navbar">
        <div className="navbar-brand">
          <DigitalWalletIcon className="navbar-icon" />
          <span className="navbar-title">Cozy Wallet</span>
        </div>
        <div className="navbar-actions">
          <UserIcon className="navbar-icon" />
        </div>
      </nav>
      <div className="content">{children}</div>
    </>
  );
}
