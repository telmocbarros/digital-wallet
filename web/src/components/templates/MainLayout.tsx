import { type ReactNode, useState } from 'react';
import './MainLayout.css';
import Navbar from '../organisms/Navbar';
import SlideMenu from '../organisms/SlideMenu';

interface MainLayoutProps {
  children: ReactNode;
}

export default function MainLayout({ children }: MainLayoutProps) {
  const [displayAppMenu, setDisplayAppMenu] = useState(false);
  const [displaySettingsMenu, setDisplaySettingsMenu] = useState(false);

  function handleAppIconClick() {
    setDisplayAppMenu(!displayAppMenu);
    setDisplaySettingsMenu(false); // Close other menu
  }

  function handleMenuIconClick() {
    setDisplaySettingsMenu(!displaySettingsMenu);
    setDisplayAppMenu(false); // Close other menu
  }

  return (
    <>
      <Navbar
        onAppIconClick={handleAppIconClick}
        onHamburgerIconClick={handleMenuIconClick}
      />

      <SlideMenu isOpen={displaySettingsMenu} side="right">
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
