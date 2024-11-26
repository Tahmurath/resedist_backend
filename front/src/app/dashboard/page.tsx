"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import React, { useState,useEffect } from "react";
import { toast } from "@/hooks/use-toast"
// import { ToastAction } from "@/components/ui/toast"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
// import Link from 'next/link'
// import { buttonVariants } from "@/components/ui/button"


export default function Home() {


    const [advice, setAdvice] = useState("");
    const [count, setCount] = useState<number>(0);
    const [error, setError] = useState<any>(null);

    async function getAdvice() {
        //setLoading(true);
        setError(null); // پاک کردن خطای قبلی

        try {
            const res = await fetch("https://api.adviceslip.com/advice?" + new Date().getTime());
            if (!res.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await res.json();
            setAdvice(data.slip.advice);
            toastaction({messagetxt:data.slip.advice})
            setCount((c) => c + 1);
        } catch (error) {
            console.error("Error fetching data:", error);
            setError("خطا در دریافت داده‌ها. لطفاً بعداً تلاش کنید.");
        } finally {
            //setLoading(false);
        }
    }
    async function getUser() {
        //setLoading(true);
        setError(null); // پاک کردن خطای قبلی

        const token = getTokenFromCookie()

        try {
            const res = await fetch("http://localhost:8080/api/v1/auth/user?" + new Date().getTime(),
            {
              method: 'GET',
              headers: {
              'Content-Type': 'application/json',
              'Authorization': `Bearer ${token}`, // ست کردن هدر Authorization
            }
          }
          );
            if (!res.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await res.json();
          
            // toastaction({messagetxt:data.message})
            toastaction({messagetxt:data.user.Email})
            console.info(data)
        } catch (error) {
            console.error("Error fetching data:", error);
            setError("خطا در دریافت داده‌ها. لطفاً بعداً تلاش کنید.");
        } finally {
            //setLoading(false);
        }
    }
    function getJWT(){
      const token = getTokenFromCookie()
    }

    const Message = (count:number):any => {
        return (
            <p>You read <strong>{count}</strong></p>
        );
    }

    useEffect(() => {
        //getAdvice();
    }, []);

    function toastaction({messagetxt}:{messagetxt: string}) {
      toast({
        title: "You submitted the following values:",
        description: (
          <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
            <code className="text-white">{messagetxt}</code>
          </pre>
        ),
      })
    }
    
    const FormSchema = z.object({
      username: z.string().min(2, {
        message: "Username must be at least 2 characters.",
      }),
      password: z.string().min(2, {
        message: "Password must be at least 2 characters.",
      }),
    })
    
    // type InputFormProps = {
    //     username: string;
    //     password: string;
    // };
    // export function InputForm({ username, password }: InputFormProps) {
    //<Button variant="outline" onClick={toastaction({messagetxt:"23423423"})}>toastaction</Button>
    
    
    const saveTokenToCookie = (token) => {
      document.cookie = `Bearer=${token}; max-age=${7 * 24 * 60 * 60}`;
    };

    const getTokenFromCookie = () => {
      const cookies = document.cookie.split('; ');
      const tokenCookie = cookies.find((row) => row.startsWith('Bearer='));
      //console.info(tokenCookie.split('=')[1])
      return tokenCookie ? tokenCookie.split('=')[1] : null;
    };
    
    
    function InputForm({username,password}:{username: string,password: string}) {
    
      const form = useForm<z.infer<typeof FormSchema>>({
        resolver: zodResolver(FormSchema),
        defaultValues: {
          username: username,
          password: password,
        },
      })
    
      function onSubmit(data: z.infer<typeof FormSchema>) {
        //alert(JSON.stringify(data, null, 2))
    
        const token = getTokenFromCookie()
          fetch("http://127.0.0.1:8080/api/v1/auth/login", {
              method: "POST",
              
              body: JSON.stringify({
                  email: data.username,
                  password: data.password
              }),
              headers: {
                  "Content-type": "application/json; charset=UTF-8",
                  //'Authorization': `Bearer ${token}`,
              }
          })
              .then((response) => response.json())
              .then((json) => 
                {
                  console.log(json)
                  saveTokenToCookie(json.token)
                  toast({
                    title: "You submitted the following values:",
                    description: (
                      <pre className="mt-2 w rounded-md bg-slate-950 p-4">
                        <code className="text-white">{json.token}</code>
                      </pre>
                    ),
                  })
                });
    
        toast({
          title: "You submitted the following values:",
          description: (
            <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
              <code className="text-white">{JSON.stringify(data, null, 2)}</code>
            </pre>
          ),
        })
      }
    
      return (
        <Form {...form}>
          <form onSubmit={form.handleSubmit(onSubmit)} className="w-2/3 space-y-6">
            <FormField
              control={form.control}
              name="username"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input placeholder="shadcn" {...field}
                           // type="text"
                           // value={inputValue}
                           // onChange={(e: React.ChangeEvent<HTMLInputElement>) => setInputValue(e.target.value)}
                    />
                  </FormControl>
                  <FormDescription>
                    This is your public display name.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>password</FormLabel>
                  <FormControl>
                    <Input placeholder="shadcn" {...field}
                           // type="text"
                           // value={inputValue2}
                           // onChange={(e: React.ChangeEvent<HTMLInputElement>) => setInputValue2(e.target.value)}
                    />
                  </FormControl>
                  <FormDescription>
                    This is your password.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button type="submit">Submit</Button>
          </form>
        </Form>
      )
    }

  return (
      <div>
          <h1>Dashboard</h1>
          {Message(count)}
          <h1>{advice}</h1>
          <Button onClick={getAdvice}>get Advice</Button>
          <Button onClick={getUser}>getJWT</Button>
          <InputForm password={""} username={''}></InputForm>
      </div>
  );
}


