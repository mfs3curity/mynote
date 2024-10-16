مشروع بسيط تستطيع استخدامه على جهازك وبناء واجهة front end 
بعد تحميل المشروع على جهازك تشغيله بالامر 
```bash
go run cmd/main.go

```
#### الهدف من المشروع انك تتعلم بناء واجهات أو تجعله كمكتبة خاصه بجهازك  تكتب فيها شروحاتك الخاصه




ساقوم بكتابة نقاط النهاية

 
اليوزر الافتراضي والرقم السري `admin`
```bash

curl http://localhost:1337/api/auth/login -d '{"username":"admin","password":"admin"}' 
```
انشاء قسم جديد

```bash
curl http://localhost:1337/api/section/create -d '{"name":"golang"}' -H "Authorization: Bearer <Token>"
```

عرض القسم وعدد البوستات لو كانت موجوده مثلا لو هناك 100 بوست وترغب بعرض عشرهه كل مره 
```bash
curl http://localhost:1337/api/section/name/vulnerabilities?limit=10&offset=0
```

انشاء بوست جديد

```
curl http://localhost:1337/api/post/create -d '{"title":"title1","description":"i am content", "section_name":"vulnerabilities"}' -H "Authorization: Bearer <Token>

```

عرض البوست 
```
curl http://localhost:1337/api/post/id/1
```

تحميل الصور 

```
curl -X POST http://localhost:1337/api/post/upload -F "file=@/home/mf/Pictures/s.jpg" -H "Authorization: Bearer <Token>

```
اضف نقطة النهاية تحميل الصور وقت اضافة بوست جديد لكي يقوم بتحميل الصورة ويتم ارجاع مسار الصوره وادراجها داخل البوست 

