import HomeIcon from './HomeIcon';
import InvoiceIcon from './InvoiceIcon';
import SettingsIcon from './SettingsIcon';
import Transactions from './Transactions';
import DefaultIcon from './DefaultIcon';

export default function Icon({
  name,
  ...props
}: React.SVGProps<SVGSVGElement>) {
  switch (name) {
    case 'home':
      return <HomeIcon {...props} />;
    case 'invoice':
      return <InvoiceIcon {...props} />;
    case 'settings':
      return <SettingsIcon {...props} />;
    case 'transactions':
      return <Transactions {...props} />;
    default:
      return <DefaultIcon {...props} />;
  }
}

// TODO: Add all the other icons here
