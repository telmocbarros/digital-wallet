export default function CardMolecule({
  title = '',
  content = '',
  className = '',
}) {
  return (
    <div
      className={`w-56 h-56 bg-amber-50 shadow-xl p-6 m-6 rounded-xl ${className}`}
    >
      <h2>{title}</h2>
      <p>{content}</p>
      <p>
        Lorem ipsum dolor sit, amet consectetur adipisicing elit. Iste, iure
        sint hic quisquam nesciunt est non harum id, inventore laboriosam
        molestias porro odio! Optio, obcaecati dolores inventore rem vel nulla!
      </p>
    </div>
  );
}
