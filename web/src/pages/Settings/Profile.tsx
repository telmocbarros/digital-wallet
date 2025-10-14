import { useState } from 'react';
import Button from '../../components/atoms/Button';

export default function Profile({ className }: { className?: string }) {
  const [displayName, setDisplayName] = useState('John Doe');
  const [bio, setBio] = useState(
    'Financial enthusiast and digital wallet user.'
  );
  const [location, setLocation] = useState('San Francisco, CA');
  const [website, setWebsite] = useState('https://johndoe.com');
  const [twitter, setTwitter] = useState('@johndoe');
  const [linkedin, setLinkedin] = useState('linkedin.com/in/johndoe');
  const [profileVisibility, setProfileVisibility] = useState<
    'public' | 'private'
  >('public');
  const [showTransactionHistory, setShowTransactionHistory] = useState(false);

  const handleSave = () => {
    console.log('Saving profile settings:', {
      displayName,
      bio,
      location,
      website,
      twitter,
      linkedin,
      profileVisibility,
      showTransactionHistory,
    });
    // TODO: Implement API call to save settings
  };

  const handleCancel = () => {
    console.log('Cancelling changes');
    // TODO: Reset to previous values
  };

  const handlePhotoUpload = () => {
    console.log('Opening file picker for profile photo...');
    // TODO: Implement photo upload
  };

  return (
    <div className={`max-w-4xl ${className}`}>
      {/* Header */}
      <div className="mb-8">
        <h2 className="text-2xl font-semibold text-gray-900 mb-2">Profile</h2>
        <p className="text-gray-600">
          Manage your public profile information and how others see you.
        </p>
      </div>

      {/* Profile Photo Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-4">
          Profile Photo
        </h3>
        <div className="flex items-start gap-6">
          <div className="relative">
            <div className="w-24 h-24 rounded-full bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white text-3xl font-semibold">
              JD
            </div>
            <button
              onClick={handlePhotoUpload}
              className="absolute bottom-0 right-0 w-8 h-8 bg-white border-2 border-gray-200 rounded-full flex items-center justify-center hover:bg-gray-50 shadow-sm"
            >
              <svg
                className="w-4 h-4 text-gray-600"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M3 9a2 2 0 012-2h.93a2 2 0 001.664-.89l.812-1.22A2 2 0 0110.07 4h3.86a2 2 0 011.664.89l.812 1.22A2 2 0 0018.07 7H19a2 2 0 012 2v9a2 2 0 01-2 2H5a2 2 0 01-2-2V9z"
                />
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M15 13a3 3 0 11-6 0 3 3 0 016 0z"
                />
              </svg>
            </button>
          </div>
          <div className="flex-1">
            <button
              onClick={handlePhotoUpload}
              className="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 mb-2"
            >
              Upload new photo
            </button>
            <p className="text-xs text-gray-500">
              JPG, PNG or GIF. Max size 2MB. Recommended: 400x400px.
            </p>
          </div>
        </div>
      </div>

      {/* Public Profile Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-4">
          Public Profile
        </h3>

        {/* Display Name */}
        <div className="mb-4">
          <label
            htmlFor="displayName"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Display name
          </label>
          <input
            id="displayName"
            type="text"
            value={displayName}
            onChange={(e) => setDisplayName(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Enter your display name"
          />
          <p className="text-xs text-gray-500 mt-1">
            This is how your name will appear to other users.
          </p>
        </div>

        {/* Bio */}
        <div className="mb-4">
          <label
            htmlFor="bio"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Bio
          </label>
          <textarea
            id="bio"
            value={bio}
            onChange={(e) => setBio(e.target.value)}
            rows={4}
            maxLength={160}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent resize-none"
            placeholder="Write a short bio about yourself..."
          />
          <div className="flex justify-between mt-1">
            <p className="text-xs text-gray-500">
              Brief description for your profile.
            </p>
            <p className="text-xs text-gray-500">{bio.length}/160</p>
          </div>
        </div>

        {/* Location */}
        <div className="mb-4">
          <label
            htmlFor="location"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Location
          </label>
          <input
            id="location"
            type="text"
            value={location}
            onChange={(e) => setLocation(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="City, Country"
          />
        </div>

        {/* Website */}
        <div className="mb-4">
          <label
            htmlFor="website"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Website
          </label>
          <input
            id="website"
            type="url"
            value={website}
            onChange={(e) => setWebsite(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="https://yourwebsite.com"
          />
        </div>
      </div>

      {/* Social Links Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-4">Social Links</h3>

        {/* Twitter */}
        <div className="mb-4">
          <label
            htmlFor="twitter"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            Twitter / X
          </label>
          <div className="flex items-center">
            <span className="inline-flex items-center px-3 py-2 border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm rounded-l-lg">
              <svg className="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                <path d="M18.244 2.25h3.308l-7.227 8.26 8.502 11.24H16.17l-5.214-6.817L4.99 21.75H1.68l7.73-8.835L1.254 2.25H8.08l4.713 6.231zm-1.161 17.52h1.833L7.084 4.126H5.117z" />
              </svg>
            </span>
            <input
              id="twitter"
              type="text"
              value={twitter}
              onChange={(e) => setTwitter(e.target.value)}
              className="flex-1 px-3 py-2 border border-gray-300 rounded-r-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="@username"
            />
          </div>
        </div>

        {/* LinkedIn */}
        <div className="mb-4">
          <label
            htmlFor="linkedin"
            className="block text-sm font-medium text-gray-700 mb-2"
          >
            LinkedIn
          </label>
          <div className="flex items-center">
            <span className="inline-flex items-center px-3 py-2 border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm rounded-l-lg">
              <svg className="w-4 h-4" fill="currentColor" viewBox="0 0 24 24">
                <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z" />
              </svg>
            </span>
            <input
              id="linkedin"
              type="text"
              value={linkedin}
              onChange={(e) => setLinkedin(e.target.value)}
              className="flex-1 px-3 py-2 border border-gray-300 rounded-r-lg text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="linkedin.com/in/username"
            />
          </div>
        </div>
      </div>

      {/* Privacy Settings Section */}
      <div className="mb-8 pb-8">
        <h3 className="text-sm font-medium text-gray-900 mb-4">
          Privacy Settings
        </h3>

        {/* Profile Visibility */}
        <div className="mb-4">
          <label className="block text-sm font-medium text-gray-700 mb-3">
            Profile visibility
          </label>
          <div className="space-y-3">
            <label className="flex items-start cursor-pointer">
              <input
                type="radio"
                name="visibility"
                value="public"
                checked={profileVisibility === 'public'}
                onChange={(e) =>
                  setProfileVisibility(e.target.value as 'public' | 'private')
                }
                className="mt-1 h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300"
              />
              <div className="ml-3">
                <span className="text-sm font-medium text-gray-900">
                  Public
                </span>
                <p className="text-xs text-gray-500">
                  Anyone can view your profile and activity.
                </p>
              </div>
            </label>
            <label className="flex items-start cursor-pointer">
              <input
                type="radio"
                name="visibility"
                value="private"
                checked={profileVisibility === 'private'}
                onChange={(e) =>
                  setProfileVisibility(e.target.value as 'public' | 'private')
                }
                className="mt-1 h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300"
              />
              <div className="ml-3">
                <span className="text-sm font-medium text-gray-900">
                  Private
                </span>
                <p className="text-xs text-gray-500">
                  Only you can see your profile information.
                </p>
              </div>
            </label>
          </div>
        </div>

        {/* Show Transaction History */}
        <div className="mt-6">
          <label className="flex items-start cursor-pointer">
            <input
              type="checkbox"
              checked={showTransactionHistory}
              onChange={(e) => setShowTransactionHistory(e.target.checked)}
              className="mt-1 h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
            />
            <div className="ml-3">
              <span className="text-sm font-medium text-gray-900">
                Show transaction history
              </span>
              <p className="text-xs text-gray-500">
                Allow others to see your public transaction history (amounts are
                always hidden).
              </p>
            </div>
          </label>
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
