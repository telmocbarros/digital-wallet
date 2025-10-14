import { useState } from 'react';
import './MainLayout.css';
import Navbar from '../organisms/Navbar';
import SlideMenu from '../organisms/SlideMenu';
import { Outlet } from 'react-router';
import SideBar from '../organisms/SideBar';

export default function MainLayout() {
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
      <div className="main-layout-container">
        <SideBar />
        <div className="content">
          <h1>content</h1>
          <Outlet />
        </div>
      </div>

      <SlideMenu isOpen={displaySettingsMenu} side="right">
        <ul>
          <li>Profile</li>
          <li>Settings</li>
          <li>Logout</li>
        </ul>
      </SlideMenu>
    </>
  );
}
