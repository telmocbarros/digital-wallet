import { useState } from 'react';
import Button from '../../components/atoms/Button';

type SessionDevice = {
  id: string;
  device: string;
  location: string;
  lastActive: string;
  isCurrent: boolean;
};

export default function Security({ className }: { className?: string }) {
  const [currentPassword, setCurrentPassword] = useState('');
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [twoFactorEnabled, setTwoFactorEnabled] = useState(false);
  const [securityAlerts, setSecurityAlerts] = useState(true);
  const [loginNotifications, setLoginNotifications] = useState(true);

  const [activeSessions] = useState<SessionDevice[]>([
    {
      id: '1',
      device: 'Chrome on MacBook Pro',
      location: 'San Francisco, CA',
      lastActive: '2 minutes ago',
      isCurrent: true,
    },
    {
      id: '2',
      device: 'Safari on iPhone 15',
      location: 'San Francisco, CA',
      lastActive: '2 hours ago',
      isCurrent: false,
    },
    {
      id: '3',
      device: 'Chrome on Windows',
      location: 'New York, NY',
      lastActive: '3 days ago',
      isCurrent: false,
    },
  ]);

  const handlePasswordChange = () => {
    if (newPassword !== confirmPassword) {
      alert('Passwords do not match!');
      return;
    }
    console.log('Changing password...');
    // TODO: Implement password change API call
  };

  const handleEnable2FA = () => {
    console.log('Enabling 2FA...');
    setTwoFactorEnabled(true);
    // TODO: Implement 2FA setup flow
  };

  const handleDisable2FA = () => {
    const confirmed = window.confirm(
      'Are you sure you want to disable two-factor authentication? This will make your account less secure.'
    );
    if (confirmed) {
      console.log('Disabling 2FA...');
      setTwoFactorEnabled(false);
      // TODO: Implement 2FA disable API call
    }
  };

  const handleRevokeSession = (sessionId: string) => {
    console.log('Revoking session:', sessionId);
    // TODO: Implement session revocation
  };

  const handleRevokeAllSessions = () => {
    const confirmed = window.confirm(
      'This will sign you out from all devices except this one. Continue?'
    );
    if (confirmed) {
      console.log('Revoking all sessions...');
      // TODO: Implement revoke all sessions
    }
  };

  return (
    <div className={`max-w-4xl ${className}`}>
      {/* Header */}
      <div className="mb-8">
        <h2 className="text-2xl font-semibold text-gray-900 mb-2">Security</h2>
        <p className="text-gray-600">
          Manage your account security and authentication settings.
        </p>
      </div>

      {/* Password Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-4">
          Change Password
        </h3>

        <div className="space-y-4 max-w-md">
          {/* Current Password */}
          <div>
            <label
              htmlFor="currentPassword"
              className="block text-sm font-medium text-gray-700 mb-2"
            >
              Current password
            </label>
            <input
              id="currentPassword"
              type="password"
              value={currentPassword}
              onChange={(e) => setCurrentPassword(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="Enter current password"
            />
          </div>

          {/* New Password */}
          <div>
            <label
              htmlFor="newPassword"
              className="block text-sm font-medium text-gray-700 mb-2"
            >
              New password
            </label>
            <input
              id="newPassword"
              type="password"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="Enter new password"
            />
            <p className="text-xs text-gray-500 mt-1">
              Must be at least 8 characters with letters, numbers, and symbols.
            </p>
          </div>

          {/* Confirm Password */}
          <div>
            <label
              htmlFor="confirmPassword"
              className="block text-sm font-medium text-gray-700 mb-2"
            >
              Confirm new password
            </label>
            <input
              id="confirmPassword"
              type="password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="Confirm new password"
            />
          </div>

          <Button
            label="Update Password"
            onClick={handlePasswordChange}
            disabled={!currentPassword || !newPassword || !confirmPassword}
          />
        </div>
      </div>

      {/* Two-Factor Authentication Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-4">
          Two-Factor Authentication (2FA)
        </h3>

        {!twoFactorEnabled ? (
          <div className="bg-amber-50 rounded-lg p-4 border border-amber-200">
            <div className="flex items-start">
              <svg
                className="w-5 h-5 text-amber-600 mr-3 mt-0.5"
                fill="currentColor"
                viewBox="0 0 20 20"
              >
                <path
                  fillRule="evenodd"
                  d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
                  clipRule="evenodd"
                />
              </svg>
              <div className="flex-1">
                <h4 className="text-sm font-medium text-amber-900 mb-1">
                  2FA is not enabled
                </h4>
                <p className="text-sm text-amber-700 mb-3">
                  Add an extra layer of security to your account by requiring
                  more than just a password to sign in.
                </p>
                <button
                  onClick={handleEnable2FA}
                  className="px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-lg hover:bg-blue-700"
                >
                  Enable Two-Factor Authentication
                </button>
              </div>
            </div>
          </div>
        ) : (
          <div className="bg-green-50 rounded-lg p-4 border border-green-200">
            <div className="flex items-start justify-between">
              <div className="flex items-start">
                <svg
                  className="w-5 h-5 text-green-600 mr-3 mt-0.5"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    fillRule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                    clipRule="evenodd"
                  />
                </svg>
                <div>
                  <h4 className="text-sm font-medium text-green-900 mb-1">
                    2FA is enabled
                  </h4>
                  <p className="text-sm text-green-700 mb-2">
                    Your account is protected with two-factor authentication.
                  </p>
                  <p className="text-xs text-green-600">
                    Method: Authenticator App (Google Authenticator, Authy)
                  </p>
                </div>
              </div>
              <button
                onClick={handleDisable2FA}
                className="px-3 py-1.5 text-xs font-medium text-red-700 bg-white border border-red-300 rounded-lg hover:bg-red-50"
              >
                Disable
              </button>
            </div>
          </div>
        )}
      </div>

      {/* Active Sessions Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <div className="flex items-center justify-between mb-4">
          <h3 className="text-sm font-medium text-gray-900">Active Sessions</h3>
          <button
            onClick={handleRevokeAllSessions}
            className="text-sm text-red-600 hover:text-red-700 font-medium"
          >
            Sign out all devices
          </button>
        </div>

        <p className="text-sm text-gray-600 mb-4">
          Manage devices where you're currently signed in.
        </p>

        <div className="space-y-3">
          {activeSessions.map((session) => (
            <div
              key={session.id}
              className="flex items-start justify-between p-4 bg-gray-50 rounded-lg border border-gray-200"
            >
              <div className="flex items-start">
                <svg
                  className="w-5 h-5 text-gray-500 mr-3 mt-0.5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                  />
                </svg>
                <div>
                  <div className="flex items-center gap-2">
                    <h4 className="text-sm font-medium text-gray-900">
                      {session.device}
                    </h4>
                    {session.isCurrent && (
                      <span className="inline-flex items-center px-2 py-0.5 text-xs font-medium text-green-700 bg-green-100 rounded">
                        Current
                      </span>
                    )}
                  </div>
                  <p className="text-xs text-gray-600 mt-1">
                    {session.location}
                  </p>
                  <p className="text-xs text-gray-500 mt-0.5">
                    Last active: {session.lastActive}
                  </p>
                </div>
              </div>
              {!session.isCurrent && (
                <button
                  onClick={() => handleRevokeSession(session.id)}
                  className="text-sm text-red-600 hover:text-red-700 font-medium"
                >
                  Revoke
                </button>
              )}
            </div>
          ))}
        </div>
      </div>

      {/* Security Notifications Section */}
      <div className="mb-8 pb-8">
        <h3 className="text-sm font-medium text-gray-900 mb-4">
          Security Notifications
        </h3>

        <div className="space-y-4">
          {/* Security Alerts */}
          <label className="flex items-start cursor-pointer">
            <input
              type="checkbox"
              checked={securityAlerts}
              onChange={(e) => setSecurityAlerts(e.target.checked)}
              className="mt-1 h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <div className="ml-3">
              <span className="text-sm font-medium text-gray-900">
                Security alerts
              </span>
              <p className="text-xs text-gray-500">
                Receive notifications about unusual account activity, failed
                login attempts, and security issues.
              </p>
            </div>
          </label>

          {/* Login Notifications */}
          <label className="flex items-start cursor-pointer">
            <input
              type="checkbox"
              checked={loginNotifications}
              onChange={(e) => setLoginNotifications(e.target.checked)}
              className="mt-1 h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <div className="ml-3">
              <span className="text-sm font-medium text-gray-900">
                Login notifications
              </span>
              <p className="text-xs text-gray-500">
                Get notified when your account is accessed from a new device or
                location.
              </p>
            </div>
          </label>
        </div>
      </div>

      {/* Action Buttons */}
      <div className="flex justify-end gap-3 pt-6 border-t border-gray-200">
        <button className="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50">
          Cancel
        </button>
        <Button
          label="Save changes"
          onClick={() => console.log('Saving security settings...')}
        />
      </div>
    </div>
  );
}
