import Button from '../../atoms/Button';
import BluewaveBankCard from '../../cards_delete_in_the_future/BlueBank/ReactComponent/BlueCardComponent';
import GreenwaveBankCard from '../../cards_delete_in_the_future/GreenBank/ReactComponent/GreenCardComponent';
import RedwaveBankCard from '../../cards_delete_in_the_future/RedBank/ReactComponent/RedCardComponent';
import YellowwaveBankCard from '../../cards_delete_in_the_future/YellowBank/ReactComponent/YelloCardComponent';
import CardMolecule from '../../molecules/CardComponent';
import SearchBar from '../../molecules/SearchBar';

export default function BankCardsList(props: { maxHeight?: string }) {
  const { maxHeight } = props; // e.g., "400px", "50vh", "max-h-96"}

  const heightClass = maxHeight?.startsWith('max-h-') ? maxHeight : '';

  const inlineStyle =
    maxHeight && !maxHeight.startsWith('max-h-')
      ? { maxHeight: maxHeight }
      : {};
  return (
    <div>
      {/* Search Bar */}
      <SearchBar />
      {/* Add Card Button */}
      <Button label="Add New Card" onClick={() => {}} />
      {/* List of Bank Cards */}
      <div className="flex flex-row">
        <div
          className={`overflow-y-auto ${heightClass || 'max-h-[500px]'}`}
          style={inlineStyle}
        >
          <BluewaveBankCard />
          <RedwaveBankCard />
          <GreenwaveBankCard />
          <YellowwaveBankCard />
          <BluewaveBankCard />
          <RedwaveBankCard />
          <GreenwaveBankCard />
          <YellowwaveBankCard />
          <BluewaveBankCard />
          <RedwaveBankCard />
          <GreenwaveBankCard />
          <YellowwaveBankCard />
          <BluewaveBankCard />
          <RedwaveBankCard />
          <GreenwaveBankCard />
          <YellowwaveBankCard />
        </div>
        <CardMolecule
          title={'CARD SELECTED'}
          content={'MAMBO JAMBO SEMPRE RAMBO'}
          className="flex-2"
        />
      </div>
    </div>
  );
}
