import { useState } from 'react';
import Button from '../../components/atoms/Button';

type ChartStyle = 'default' | 'simplified' | 'custom';
type CookieBanner = 'default' | 'simplified' | 'none';

export default function Appearance() {
  const [brandColor, setBrandColor] = useState('#444CE7');
  const [chartStyle, setChartStyle] = useState<ChartStyle>('default');
  const [language, setLanguage] = useState('en-UK');
  const [cookieBanner, setCookieBanner] = useState<CookieBanner>('default');

  const handleSave = () => {
    console.log('Saving appearance settings:', {
      brandColor,
      chartStyle,
      language,
      cookieBanner,
    });
    // TODO: Implement API call to save settings
  };

  const handleCancel = () => {
    console.log('Cancelling changes');
    // TODO: Reset to previous values
  };

  return (
    <div className="max-w-4xl">
      {/* Header */}
      <div className="mb-8">
        <div className="flex justify-between items-start mb-2">
          <h2 className="text-2xl font-semibold text-gray-900">Appearance</h2>
          <a
            href="https://dashboard.untitledui.com"
            target="_blank"
            rel="noopener noreferrer"
            className="text-sm text-gray-500 hover:text-gray-700 flex items-center gap-1"
          >
            dashboard.untitledui.com
            <svg
              className="w-4 h-4"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
              />
            </svg>
          </a>
        </div>
        <p className="text-gray-600">
          Change how your public dashboard looks and feels.
        </p>
      </div>

      {/* Brand Color Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-1">Brand color</h3>
        <p className="text-sm text-gray-600 mb-4">
          Select or customize your brand color.
        </p>
        <div className="flex items-center gap-3">
          <div
            className="w-10 h-10 rounded-lg border border-gray-200 cursor-pointer"
            style={{ backgroundColor: brandColor }}
            onClick={() => document.getElementById('colorPicker')?.click()}
          />
          <input
            id="colorPicker"
            type="color"
            value={brandColor}
            onChange={(e) => setBrandColor(e.target.value)}
            className="hidden"
          />
          <input
            type="text"
            value={brandColor}
            onChange={(e) => setBrandColor(e.target.value)}
            className="px-3 py-2 border border-gray-300 rounded-lg text-sm w-32"
            placeholder="#444CE7"
          />
        </div>
      </div>

      {/* Dashboard Charts Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-1">
          Dashboard charts
        </h3>
        <p className="text-sm text-gray-600 mb-2">How charts are displayed.</p>
        <a href="#" className="text-sm text-blue-600 hover:text-blue-700 mb-4 inline-block">
          View examples
        </a>

        <div className="grid grid-cols-3 gap-4">
          {/* Default Option */}
          <div
            onClick={() => setChartStyle('default')}
            className={`relative border-2 rounded-lg p-4 cursor-pointer transition-all ${
              chartStyle === 'default'
                ? 'border-blue-600 bg-blue-50'
                : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            {chartStyle === 'default' && (
              <div className="absolute top-2 right-2 w-5 h-5 bg-blue-600 rounded-full flex items-center justify-center">
                <svg className="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                </svg>
              </div>
            )}
            <div className="bg-white border border-gray-200 rounded p-2 mb-3 h-24">
              <div className="text-xs text-gray-500 mb-2">Dashboard</div>
              <svg className="w-full h-16" viewBox="0 0 200 60">
                <polyline
                  points="0,40 40,30 80,35 120,20 160,25 200,15"
                  fill="none"
                  stroke="#444CE7"
                  strokeWidth="2"
                />
                <polyline
                  points="0,50 40,45 80,48 120,40 160,42 200,35"
                  fill="none"
                  stroke="#9CA3AF"
                  strokeWidth="2"
                />
              </svg>
            </div>
            <h4 className="text-sm font-medium text-gray-900 mb-1">Default</h4>
            <p className="text-xs text-gray-600">Default company branding.</p>
          </div>

          {/* Simplified Option */}
          <div
            onClick={() => setChartStyle('simplified')}
            className={`relative border-2 rounded-lg p-4 cursor-pointer transition-all ${
              chartStyle === 'simplified'
                ? 'border-blue-600 bg-blue-50'
                : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            {chartStyle === 'simplified' && (
              <div className="absolute top-2 right-2 w-5 h-5 bg-blue-600 rounded-full flex items-center justify-center">
                <svg className="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                </svg>
              </div>
            )}
            <div className="bg-white border border-gray-200 rounded p-2 mb-3 h-24">
              <div className="text-xs text-gray-500 mb-2">Dashboard</div>
              <svg className="w-full h-16" viewBox="0 0 200 60">
                <polyline
                  points="0,35 50,25 100,30 150,15 200,20"
                  fill="none"
                  stroke="#6B7280"
                  strokeWidth="1.5"
                />
              </svg>
            </div>
            <h4 className="text-sm font-medium text-gray-900 mb-1">Simplified</h4>
            <p className="text-xs text-gray-600">Minimal and modern.</p>
          </div>

          {/* Custom CSS Option */}
          <div
            onClick={() => setChartStyle('custom')}
            className={`relative border-2 rounded-lg p-4 cursor-pointer transition-all ${
              chartStyle === 'custom'
                ? 'border-blue-600 bg-blue-50'
                : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            {chartStyle === 'custom' && (
              <div className="absolute top-2 right-2 w-5 h-5 bg-blue-600 rounded-full flex items-center justify-center">
                <svg className="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                </svg>
              </div>
            )}
            <div className="bg-gray-100 border border-gray-200 rounded p-2 mb-3 h-24 flex items-center justify-center">
              <div className="text-center">
                <div className="text-xs text-gray-400 mb-1">{'</> Edit CSS'}</div>
              </div>
            </div>
            <h4 className="text-sm font-medium text-gray-900 mb-1">Custom CSS</h4>
            <p className="text-xs text-gray-600">Manage styling with CSS.</p>
          </div>
        </div>
      </div>

      {/* Language Section */}
      <div className="mb-8 pb-8 border-b border-gray-200">
        <h3 className="text-sm font-medium text-gray-900 mb-1">Language</h3>
        <p className="text-sm text-gray-600 mb-4">
          Default language for public dashboard.
        </p>
        <div className="relative max-w-xs">
          <select
            value={language}
            onChange={(e) => setLanguage(e.target.value)}
            className="w-full px-3 py-2 pr-10 border border-gray-300 rounded-lg text-sm appearance-none bg-white"
          >
            <option value="en-UK">🇬🇧 English (UK)</option>
            <option value="en-US">🇺🇸 English (US)</option>
            <option value="es">🇪🇸 Español</option>
            <option value="fr">🇫🇷 Français</option>
            <option value="de">🇩🇪 Deutsch</option>
            <option value="pt">🇵🇹 Português</option>
          </select>
          <div className="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none">
            <svg className="w-4 h-4 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
            </svg>
          </div>
        </div>
      </div>

      {/* Cookie Banner Section */}
      <div className="mb-8 pb-8">
        <h3 className="text-sm font-medium text-gray-900 mb-1">Cookie banner</h3>
        <p className="text-sm text-gray-600 mb-2">
          Display cookie banners to visitors.
        </p>
        <a href="#" className="text-sm text-blue-600 hover:text-blue-700 mb-4 inline-block">
          View examples
        </a>

        <div className="grid grid-cols-3 gap-4">
          {/* Default Cookie Banner */}
          <div
            onClick={() => setCookieBanner('default')}
            className={`relative border-2 rounded-lg p-4 cursor-pointer transition-all ${
              cookieBanner === 'default'
                ? 'border-blue-600 bg-blue-50'
                : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            {cookieBanner === 'default' && (
              <div className="absolute top-2 right-2 w-5 h-5 bg-blue-600 rounded-full flex items-center justify-center">
                <svg className="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                </svg>
              </div>
            )}
            <div className="bg-white border border-gray-200 rounded p-2 mb-3 h-24 relative">
              <div className="absolute bottom-2 left-2 right-2 bg-blue-600 rounded p-2 text-xs text-white">
                Cookie banner
              </div>
            </div>
            <h4 className="text-sm font-medium text-gray-900 mb-1">Default</h4>
            <p className="text-xs text-gray-600">Cookie controls for visitors.</p>
          </div>

          {/* Simplified Cookie Banner */}
          <div
            onClick={() => setCookieBanner('simplified')}
            className={`relative border-2 rounded-lg p-4 cursor-pointer transition-all ${
              cookieBanner === 'simplified'
                ? 'border-blue-600 bg-blue-50'
                : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            {cookieBanner === 'simplified' && (
              <div className="absolute top-2 right-2 w-5 h-5 bg-blue-600 rounded-full flex items-center justify-center">
                <svg className="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                </svg>
              </div>
            )}
            <div className="bg-white border border-gray-200 rounded p-2 mb-3 h-24 relative">
              <div className="absolute bottom-2 left-2 bg-gray-800 rounded px-2 py-1 text-xs text-white">
                Simple
              </div>
            </div>
            <h4 className="text-sm font-medium text-gray-900 mb-1">Simplified</h4>
            <p className="text-xs text-gray-600">Show a simplified banner.</p>
          </div>

          {/* No Cookie Banner */}
          <div
            onClick={() => setCookieBanner('none')}
            className={`relative border-2 rounded-lg p-4 cursor-pointer transition-all ${
              cookieBanner === 'none'
                ? 'border-blue-600 bg-blue-50'
                : 'border-gray-200 hover:border-gray-300'
            }`}
          >
            {cookieBanner === 'none' && (
              <div className="absolute top-2 right-2 w-5 h-5 bg-blue-600 rounded-full flex items-center justify-center">
                <svg className="w-3 h-3 text-white" fill="currentColor" viewBox="0 0 20 20">
                  <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                </svg>
              </div>
            )}
            <div className="bg-white border border-gray-200 rounded p-2 mb-3 h-24">
              {/* Empty - no banner */}
            </div>
            <h4 className="text-sm font-medium text-gray-900 mb-1">None</h4>
            <p className="text-xs text-gray-600">Don't show any banners.</p>
          </div>
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
