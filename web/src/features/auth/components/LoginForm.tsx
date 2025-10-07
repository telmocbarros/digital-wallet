import DigitalWalletIcon from '../../../assets/DigitalWalletIcon.svg';

interface LoginFormProps {
  onSubmit: (email: string, password: string) => void;
}

export default function LoginForm({ onSubmit }: LoginFormProps) {
  function handleSubmission(e: React.FormEvent) {
    e.preventDefault();
    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);
    const email = formData.get('email') as string;
    const password = formData.get('password') as string;

    onSubmit(email, password);
  }

  return (
    <div className="login-form">
      <img
        src={DigitalWalletIcon}
        alt="digital-wallet-icon"
        width="200"
        height="200"
      />
      <h2>Welcome</h2>
      <p>Enter your details to login</p>
      <form onSubmit={handleSubmission}>
        <label htmlFor="email">Email</label>
        <input
          type="email"
          id="email"
          name="email"
          placeholder="Enter email"
          autoComplete="on"
        />
        <label htmlFor="password">Password</label>
        <input
          type="password"
          id="password"
          name="password"
          placeholder="Enter your password"
          autoComplete="on"
        />
        <div className="form-password-options">
          <div className="form-password-options-checkbox">
            <input type="checkbox" />
            <p>Remember me</p>
          </div>
          <a>Forgot your password?</a>
        </div>
        <button type="submit">Login</button>
        <div className="form-new-account">
          <p>Don't have an account?</p>
          <a>Register</a>
        </div>
      </form>
    </div>
  );
}
