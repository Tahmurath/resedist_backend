"use client"
 
import { useToast } from "@/hooks/use-toast"
import { toast } from "@/hooks/use-toast"
import { ToastAction } from "@/components/ui/toast"
import { Button } from "@/components/ui/button"
import { buttonVariants } from "@/components/ui/button"
import Link from 'next/link'

export default function Home() {
  return (
    <div>
      <h1>Home</h1>
      <hr></hr>
      <Button>Click me</Button>
      <hr></hr>
      <Button variant="outline">Button</Button>
      <hr></hr>
      <Link href="/dashboard" className={buttonVariants({ variant: "outline" })}>Dashboard</Link>
      <Button
      variant="outline"
      onClick={() => {
        toast({
          title: "Scheduled: Catch up ",
          description: "Friday, February 10, 2023 at 5:57 PM",
          action: (
            <ToastAction altText="Goto schedule to undo">Undo</ToastAction>
          ),
        })
      }}
    >
      Add to calendar
    </Button>
    </div>
  );
}
