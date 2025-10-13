import BluewaveBankCard from '../cards_delete_in_the_future/BlueBank/ReactComponent/BlueCardComponent';
import GreenwaveBankCard from '../cards_delete_in_the_future/GreenBank/ReactComponent/GreenCardComponent';
import RedwaveBankCard from '../cards_delete_in_the_future/RedBank/ReactComponent/RedCardComponent';
import YellowwaveBankCard from '../cards_delete_in_the_future/YellowBank/ReactComponent/YelloCardComponent';
import './CardsList.css';

export default function CardsList({ className }) {
  // Temporary array representing our cards
  const cards = [
    {
      id: 1,
      name: 'Bluewave Bank',
      balance: 5000,
      cardComponent: <BluewaveBankCard />,
    },
    {
      id: 2,
      name: 'Green Bank',
      balance: 3200,
      cardComponent: <GreenwaveBankCard />,
    },
    {
      id: 3,
      name: 'Red Bank',
      balance: 1500,
      cardComponent: <RedwaveBankCard />,
    },
    {
      id: 4,
      name: 'Yellow Bank',
      balance: 5000,
      cardComponent: <YellowwaveBankCard />,
    },
    {
      id: 5,
      name: 'Bluewave Bank',
      balance: 3200,
      cardComponent: <BluewaveBankCard />,
    },
    {
      id: 6,
      name: 'Green Bank',
      balance: 1500,
      cardComponent: <GreenwaveBankCard />,
    },
    {
      id: 7,
      name: 'Red Bank',
      balance: 5000,
      cardComponent: <RedwaveBankCard />,
    },
    {
      id: 8,
      name: 'Yellow Bank',
      balance: 3200,
      cardComponent: <YellowwaveBankCard />,
    },
    {
      id: 9,
      name: 'Bluewave Bank',
      balance: 1500,
      cardComponent: <BluewaveBankCard />,
    },
    {
      id: 10,
      name: 'Green Bank',
      balance: 1500,
      cardComponent: <GreenwaveBankCard />,
    },
    {
      id: 11,
      name: 'Red Bank',
      balance: 5000,
      cardComponent: <RedwaveBankCard />,
    },
  ];

  const CARD_OFFSET = 60; // Vertical spacing between cards (in pixels)

  return (
    <div
      className={`flex flex-col relative w-full items-center h-[500px] p-4 scroll-smooth overflow-y-auto scrollbar-hidden ${className}`}
    >
      {cards.slice(0, 8).map((card, index) => (
        <div
          key={card.id}
          className="absolute top-0  transition-all duration-300 ease-out hover:scale-105"
          style={{
            transform: `translateY(${index * CARD_OFFSET}px)`,
            zIndex: 10 + index, // First card has highest z-index
          }}
        >
          {card.cardComponent}
        </div>
      ))}
    </div>
  );
}
