import { useTheme } from '../../../hooks/useTheme';

const Sun = (props: React.SVGProps<SVGSVGElement>) => (
  <svg
    version="1.0"
    xmlns="http://www.w3.org/2000/svg"
    width="64.000000pt"
    height="64.000000pt"
    viewBox="0 0 64.000000 64.000000"
    preserveAspectRatio="xMidYMid meet"
    {...props}
  >
    <g
      transform="translate(0.000000,64.000000) scale(0.100000,-0.100000)"
      fill="#FFEA00"
      stroke="none"
    >
      <path
        d="M302 628 c-24 -24 -11 -88 18 -88 19 0 30 19 30 54 0 21 -17 46 -30
46 -3 0 -11 -5 -18 -12z"
      />
      <path d="M90 532 c0 -23 40 -62 63 -62 25 0 21 33 -6 58 -28 27 -57 29 -57 4z" />
      <path
        d="M492 527 c-27 -28 -29 -57 -4 -57 23 0 62 40 62 63 0 25 -33 21 -58
-6z"
      />
      <path
        d="M242 474 c-110 -56 -124 -216 -25 -291 45 -35 125 -43 177 -18 114
54 129 216 29 292 -45 34 -131 42 -181 17z m148 -63 c14 -10 31 -36 39 -58 13
-36 12 -43 -5 -79 -27 -55 -81 -81 -134 -65 -72 21 -105 93 -73 159 33 68 112
88 173 43z"
      />
      <path d="M10 335 c-25 -30 26 -57 74 -39 21 8 21 40 0 48 -27 11 -61 6 -74 -9z" />
      <path
        d="M553 343 c-18 -7 -16 -40 3 -47 48 -18 99 9 74 39 -12 14 -52 19 -77
8z"
      />
      <path
        d="M112 147 c-27 -28 -29 -57 -4 -57 23 0 62 40 62 63 0 25 -33 21 -58
-6z"
      />
      <path d="M470 152 c0 -23 40 -62 63 -62 25 0 21 33 -6 58 -28 27 -57 29 -57 4z" />
      <path
        d="M296 84 c-11 -27 -6 -61 9 -74 30 -25 57 26 39 74 -8 21 -40 21 -48
0z"
      />
    </g>
  </svg>
);

const Moon = (props: React.SVGProps<SVGSVGElement>) => (
  <svg
    version="1.0"
    xmlns="http://www.w3.org/2000/svg"
    width="64.000000pt"
    height="64.000000pt"
    viewBox="0 0 64.000000 64.000000"
    preserveAspectRatio="xMidYMid meet"
    {...props}
  >
    <g
      transform="translate(0.000000,64.000000) scale(0.100000,-0.100000)"
      fill="#FFFFFF"
      stroke="none"
    >
      <path
        d="M225 526 c-60 -28 -87 -56 -114 -116 -36 -79 -19 -183 42 -249 33
-36 115 -71 167 -71 52 0 134 35 167 71 34 38 63 110 63 162 0 40 -3 47 -20
47 -12 0 -20 -7 -20 -16 0 -9 -15 -31 -34 -50 -30 -30 -40 -34 -85 -34 -44 0
-55 4 -86 35 -31 31 -35 42 -35 86 0 45 4 55 34 85 19 19 41 34 50 34 9 0 16
8 16 20 0 17 -7 20 -47 20 -29 0 -68 -9 -98 -24z m29 -63 c-30 -58 -30 -88 -1
-143 16 -30 37 -51 67 -67 55 -29 85 -29 143 1 26 13 47 19 47 14 0 -31 -59
-101 -105 -126 -39 -20 -131 -20 -170 0 -70 37 -108 100 -108 178 0 61 17 102
60 141 67 63 98 63 67 2z"
      />
    </g>
  </svg>
);

function DarkModeToggleIcon(props: React.SVGProps<SVGSVGElement>) {
  const { theme, toggleTheme } = useTheme();

  return theme === 'dark' ? (
    <Moon onClick={toggleTheme} {...props} />
  ) : (
    <Sun onClick={toggleTheme} {...props} />
  );
}

export default DarkModeToggleIcon;
