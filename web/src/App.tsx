import { useState, useEffect } from 'react';
function App() {
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080');
        const data: { message: string } = await response.json();
        setText(data.message);
      } catch (error: Error | unknown) {
        console.log(error);
        setText('Failed to fetch data');
      }
    };
    fetchData();
  }, []);
  const [text, setText] = useState('');

  return (
    <>
      <h1>{text}</h1>
    </>
  );
}

export default App;
