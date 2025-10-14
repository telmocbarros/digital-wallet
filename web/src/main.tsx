import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import './index.css';
import './api/middleware/auth.ts';
import App from './App.tsx';

import { createBrowserRouter, RouterProvider } from 'react-router';
import MainLayout from './components/templates/MainLayout.tsx';
import DashboardPage from './pages/DashboardPage.tsx';
import TransactionsPage from './pages/TransactionsPage.tsx';
import BankCardsPage from './pages/BankCardsPage.tsx';
import SettingsPage from './pages/SettingsPage.tsx';

const router = createBrowserRouter([
  {
    element: <MainLayout />,
    children: [
      {
        path: '/',
        element: <App />,
      },
      {
        path: '/dashboard',
        element: <DashboardPage />,
      },
      {
        path: '/transactions',
        element: <TransactionsPage />,
      },
      {
        path: '/cards',
        element: <BankCardsPage />,
      },
      {
        path: '/settings',
        element: <SettingsPage />,
      },
    ],
  },
]);

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>
);
