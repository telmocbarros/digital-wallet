import { useState } from 'react';
import Button from '../../components/atoms/Button';

export default function Account({ className }: { className?: string }) {
  const [name, setName] = useState('John Doe');
  const [email, setEmail] = useState('john.doe@example.com');
  const [phone, setPhone] = useState('+1 (555) 123-4567');
  const [username, setUsername] = useState('johndoe');
  const [isVerified] = useState(true);

  const handleSave = () => {
    console.log('Saving account settings:', {
      name,
      email,
      phone,
      username,
    });
    // TODO: Implement API call to save settings
  };

  const handleCancel = () => {
    console.log('Cancelling changes');
    // TODO: Reset to previous values
  };

  const handleDeleteAccount = () => {
    const confirmed = window.confirm(
      'Are you sure you want to delete your account? This action cannot be undone.'
    );
    if (confirmed) {
      console.log('Deleting account...');
      // TODO: Implement account deletion
    }
  };

  return (
    <div className={`max-w-4xl ${className}`}>
      {/* Header */}
      <div className="mb-8">
        <h2 className="text-2xl font-semibold text-gray-900 mb-2">Account</h2>
        <p className="text-gray-600">
          Manage your account information and preferences.
        </p>
      </div>

      {/* Personal Information Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-4">
          Personal Information
        </h3>

        {/* Full Name */}
        <div className="mb-4">
          <label
            htmlFor="name"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Full name
          </label>
          <input
            id="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter your full name"
          />
        </div>

        {/* Username */}
        <div className="mb-4">
          <label
            htmlFor="username"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Username
          </label>
          <input
            id="username"
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter your username"
          />
          <p className="text-xs text-gray-500 mt-1">
            Your unique username for the platform.
          </p>
        </div>

        {/* Email */}
        <div className="mb-4">
          <label
            htmlFor="email"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Email address
          </label>
          <div className="flex gap-2">
            <input
              id="email"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="flex-1 px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="Enter your email"
            />
            {isVerified && (
              <span className="inline-flex items-center px-3 py-2 bg-green-50 text-green-700 text-xs font-medium rounded-lg border border-green-200">
                <svg
                  className="w-4 h-4 mr-1"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    fillRule="evenodd"
                    d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                    clipRule="evenodd"
                  />
                </svg>
                Verified
              </span>
            )}
          </div>
          {!isVerified && (
            <p className="text-xs text-amber-600 mt-1">
              Email not verified.{' '}
              <a href="#" className="text-blue-600 hover:text-blue-700">
                Resend verification email
              </a>
            </p>
          )}
        </div>

        {/* Phone */}
        <div className="mb-4">
          <label
            htmlFor="phone"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Phone number
          </label>
          <input
            id="phone"
            type="tel"
            value={phone}
            onChange={(e) => setPhone(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="+1 (555) 000-0000"
          />
          <p className="text-xs text-gray-500 mt-1">
            Used for account recovery and two-factor authentication.
          </p>
        </div>
      </div>

      {/* Account Status Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-4">
          Account Status
        </h3>

        <div className="bg-gray-50 rounded-lg p-4 border border-gray-200">
          <div className="flex items-start justify-between mb-3">
            <div>
              <h4 className="text-sm font-medium text-gray-900 mb-1">
                Account Type
              </h4>
              <p className="text-sm text-gray-600">Standard Account</p>
            </div>
            <span className="inline-flex items-center px-3 py-1 bg-blue-100 text-blue-800 text-xs font-medium rounded-full">
              Active
            </span>
          </div>

          <div className="grid grid-cols-2 gap-4 pt-3 border-t border-gray-200">
            <div>
              <p className="text-xs text-gray-500 mb-1">Member since</p>
              <p className="text-sm font-medium text-gray-900">
                January 15, 2024
              </p>
            </div>
            <div>
              <p className="text-xs text-gray-500 mb-1">Account ID</p>
              <p className="text-sm font-medium text-gray-900 font-mono">
                ACC-2024-001234
              </p>
            </div>
          </div>
        </div>

        {/* KYC Status */}
        <div className="mt-4 bg-green-50 rounded-lg p-4 border border-green-200">
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
                KYC Verified
              </h4>
              <p className="text-sm text-green-700">
                Your identity has been verified. You have full access to all
                wallet features.
              </p>
            </div>
          </div>
        </div>
      </div>

      {/* Danger Zone Section */}
      <div className="mb-8 pb-8">
        <h3 className="text-sm font-medium text-red-900 mb-4">Danger Zone</h3>

        <div className="bg-red-50 rounded-lg p-4 border-2 border-red-200">
          <div className="flex items-start justify-between">
            <div className="flex-1">
              <h4 className="text-sm font-medium text-red-900 mb-1">
                Delete Account
              </h4>
              <p className="text-sm text-red-700 mb-3">
                Permanently delete your account and all associated data. This
                action cannot be undone.
              </p>
              <ul className="text-xs text-red-600 list-disc list-inside space-y-1">
                <li>All wallet balances must be zero</li>
                <li>All pending transactions must be completed</li>
                <li>This will remove all your personal data</li>
              </ul>
            </div>
          </div>
          <button
            onClick={handleDeleteAccount}
            className="mt-4 px-4 py-2 text-sm font-medium text-white bg-red-600 rounded-lg hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2"
          >
            Delete Account
          </button>
        </div>
      </div>

      {/* Action Buttons */}
      <div className="flex justify-end gap-3 pt-6 border-t border-gray-200">
        <button
          onClick={handleCancel}
          className="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50"
        >
          Cancel
        </button>
        <Button label="Save changes" onClick={handleSave} />
      </div>
    </div>
  );
}
