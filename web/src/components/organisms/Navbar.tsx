import DigitalWalletIcon from '../atoms/icons/DigitalWalletIcon';
import UserIcon from '../atoms/icons/UserIcon';
import SearchBar from '../molecules/SearchBar';
import './Navbar.css';

interface NavbarProps {
  onAppIconClick: () => void;
  onProfileIconClick: () => void;
  onSearch?: (value: string) => void;
}

export default function Navbar({
  onAppIconClick,
  onProfileIconClick,
  onSearch,
}: NavbarProps) {
  return (
    <nav className="main-navbar">
      <div className="navbar-brand">
        <DigitalWalletIcon onClick={onAppIconClick} className="navbar-icon" />
        <span className="navbar-title">Cozy Wallet</span>
      </div>
      <div className="navbar-actions">
        <SearchBar placeholder="Search" onSearch={onSearch} />
        <UserIcon onClick={onProfileIconClick} className="navbar-icon" />
      </div>
    </nav>
  );
}
