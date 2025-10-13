import type { JSX, SVGProps } from 'react';

function Transactions(props: SVGProps<SVGSVGElement>): JSX.Element {
  return (
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
        fill="currentColor"
        stroke="none"
      >
        <path
          d="M172 581 c-98 -8 -107 -19 -115 -143 -8 -139 3 -243 31 -270 21 -22
31 -23 200 -26 293 -5 296 -3 295 205 -2 226 -5 230 -203 236 -74 3 -168 2
-208 -2z m273 -141 c7 -11 -82 -110 -98 -110 -6 0 -21 10 -34 22 l-23 22 -39
-38 c-35 -35 -61 -41 -61 -15 0 16 82 99 98 99 9 0 25 -9 37 -20 l21 -20 34
35 c34 35 54 43 65 25z"
        />
        <path
          d="M194 85 c-12 -29 7 -35 120 -35 91 0 116 3 127 16 22 26 -7 34 -129
34 -89 0 -114 -3 -118 -15z"
        />
      </g>
    </svg>
  );
}

export default Transactions;
