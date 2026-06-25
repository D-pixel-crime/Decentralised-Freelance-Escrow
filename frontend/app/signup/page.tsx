import SignupForm from "@/components/signup-form"

export default function Signup() {
  return (
    <div className="flex-center min-h-svh w-full p-6 md:p-10"
      style={{
        background: "linear-gradient(145deg, #0a0e1a 0%, #0d1525 40%, #111b2e 100%)",
      }}
    >
      <div className="w-full max-w-sm">
        <SignupForm />
      </div>
    </div>
  )
}
