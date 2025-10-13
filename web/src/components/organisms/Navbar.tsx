import DarkModeToggleIcon from '../atoms/icons/DarkModeToggleIcon';
import DigitalWalletIcon from '../atoms/icons/DigitalWalletIcon';
import LanguageIcon from '../atoms/icons/LanguageIcon';
import UserBarIcon from '../molecules/UserBarIcon';
import NotificationBellIcon from '../atoms/icons/NotificationBellIcon';
import SearchBar from '../molecules/SearchBar';
import './Navbar.css';
import HamburgerMenuIcon from '../atoms/icons/HamburgerMenuIcon';

interface NavbarProps {
  onAppIconClick: () => void;
  onHamburgerIconClick: () => void;
  onSearch?: (value: string) => void;
}

export default function Navbar({
  onAppIconClick,
  onHamburgerIconClick,
  onSearch,
}: NavbarProps) {
  return (
    <nav className="main-navbar">
      <div className="navbar-brand">
        <DigitalWalletIcon onClick={onAppIconClick} className="navbar-icon" />
        <span className="navbar-title">Cozy Wallet</span>
      </div>
      <div className="navbar-center">
        <SearchBar
          className="navbar-desktop-only"
          placeholder="Search"
          onSearch={onSearch}
        />
      </div>
      <div className="navbar-actions">
        <HamburgerMenuIcon
          onClick={onHamburgerIconClick}
          className="navbar-icon navbar-mobile-only"
        />
        <NotificationBellIcon className="navbar-icon navbar-desktop-only" />
        <DarkModeToggleIcon className="navbar-icon navbar-desktop-only" />
        <LanguageIcon className="navbar-icon navbar-desktop-only" />
        <UserBarIcon className="navbar-desktop-only" />
      </div>
    </nav>
  );
}
