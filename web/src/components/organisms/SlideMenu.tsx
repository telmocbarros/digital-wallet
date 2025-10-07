import { type ReactNode } from 'react';
import './SlideMenu.css';

interface SlideMenuProps {
  isOpen: boolean;
  side: 'left' | 'right';
  children: ReactNode;
}

export default function SlideMenu({ isOpen, side, children }: SlideMenuProps) {
  return (
    <div className={`slide-menu slide-menu--${side} ${isOpen ? 'open' : ''}`}>
      {children}
    </div>
  );
}
