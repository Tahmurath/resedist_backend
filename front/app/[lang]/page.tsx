import { getDictionary } from './dictionaries'

type Props = {
    params: {
      lang: string;
    };
  };
  
  export default async function Page({ params }: Props) {
    const { lang } = params;
    
    const dict = await getDictionary(lang) // en
    return (


      <div>
        <h1>Language: {lang}</h1>
        {dict.products.cart}
        {/* محتوای صفحه را براساس مقدار lang تنظیم کنید */}
      </div>
    );
  }


  export async function generateStaticParams() {
    return [
      { lang: 'en' },
      { lang: 'fa' },
      // اگر زبان‌های بیشتری دارید، آن‌ها را به اینجا اضافه کنید
    ];
  }