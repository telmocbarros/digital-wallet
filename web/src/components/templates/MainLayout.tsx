import { useState } from 'react';
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
      <div className="fixed top-[60px] left-0 right-0 bottom-0 flex flex-row overflow-hidden">
        <SideBar />
        <main className="flex-1 overflow-y-hidden overflow-x-hidden p-4">
          <Outlet />
        </main>
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
