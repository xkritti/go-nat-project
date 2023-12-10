package handler

import (
	"fmt"
	"go-nat-project/database"
	"go-nat-project/models"
	"go-nat-project/utils"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

var mathDesp = map[string]string{
	"M1": `นักเรียนมีความรู้และความเข้าใจเนื้อหาคณิตศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับที่ควรพัฒนา
	นักเรียนสามารถระบุสิ่งที่โจทย์ต้องการทราบ และสิ่งที่โจทย์กำหนดได้ถูกต้องครบถ้วนเพียงบางส่วน สามารถตีความบริบทที่ให้ข้อมูลอย่างชัดเจน 
	เป็นลำดับ และบอกรายละเอียดอย่างตรงไปตรงมา ไม่มีเงื่อนไขซับซ้อน ใช้หลักการและขั้นตอนการคำนวณโดยยึดตามแบบแผนหรือหลักการที่คุ้นชิน 
	นักเรียนควรพัฒนาทักษะการวิเคราะห์ความสัมพันธ์ของปัญหาคณิตศาสตร์ 
	การใช้หลักเหตุผลควบคู่กับความคิดสร้างสรรค์ในการจินตนาการภาพจำลองสถานการณ์ซึ่งเป็นหนึ่งในเครื่องมือที่ช่วยในการแก้ปัญหาคณิตศาสตร์ 
	และควรเสริมทักษะการตีความโจทย์ที่มีความซับซ้อน วางแผนการแก้ปัญหาอย่างเป็นระบบและมีประสิทธิภาพนักเรียนมีความรู้และความเข้าใจเนื้อหาคณิตศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับที่ควรพัฒนา 
	นักเรียนสามารถระบุสิ่งที่โจทย์ต้องการทราบ และสิ่งที่โจทย์กำหนดได้ถูกต้องครบถ้วนเพียงบางส่วน สามารถตีความบริบทที่ให้ข้อมูลอย่างชัดเจน เป็นลำดับ 
	และบอกรายละเอียดอย่างตรงไปตรงมา ไม่มีเงื่อนไขซับซ้อน ใช้หลักการและขั้นตอนการคำนวณโดยยึดตามแบบแผนหรือหลักการที่คุ้นชิน นักเรียนควรพัฒนาทักษะการวิเคราะห์ความสัมพันธ์ของปัญหาคณิตศาสตร์ 
	การใช้หลักเหตุผลควบคู่กับความคิดสร้างสรรค์ในการจินตนาการภาพจำลองสถานการณ์ซึ่งเป็นหนึ่งในเครื่องมือที่ช่วยในการแก้ปัญหาคณิตศาสตร์ และควรเสริมทักษะการตีความโจทย์ที่มีความซับซ้อน วางแผนการแก้ปัญหาอย่างเป็นระบบและมีประสิทธิภาพ`,

	"M2": `นักเรียนมีความรู้และความเข้าใจเนื้อหาคณิตศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับพอใช้ นักเรียนสามารถระบุความสัมพันธ์ของปัญหาคณิตศาสตร์ สิ่งที่โจทย์ต้องการทราบ และสิ่งที่โจทย์กำหนดได้เพียงเบื้องต้น 
	โดยมีความถูกต้องและครบถ้วนเพียงบางส่วน นักเรียนสามารถใช้หลักการคิดคำนวณตามแบบแผนหรือหลักการคิดคำนวณที่คุ้นชินในการแก้ปัญหาคณิตศาสตร์ รวมถึงสามารถตีความบริบทที่มีการระบุชัดเจน และมีเงื่อนไขไม่ซับซ้อน 
	นักเรียนควรพัฒนาทักษะในการใช้เหตุผลร่วมกับความคิดสร้างสรรค์เพื่อจินตนาการภาพจำลองสถานการณ์ ในขณะเดียวกันนักเรียนควรเสริมทักษะการคิดวิเคราะห์ และเชื่อมโยงความรู้กับปัญหาคณิตศาสตร์ที่มีความซับซ้อนเพื่อที่จะสามารถแก้ปัญหาคณิตศาสตร์ได้อย่างเป็นระบบและมีประสิทธิภาพ`,

	"M3": `นักเรียนมีความรู้และความเข้าใจเนื้อหาในขอบเขตช่วงชั้นที่กำลังศึกษาระดับปานกลาง นักเรียนสามารถคิดวิเคราะห์ และระบุความสัมพันธ์ของปัญหาคณิตศาสตร์ สิ่งที่โจทย์ต้องการทราบ และสิ่งที่โจทย์กำหนดได้ในระดับมาตรฐาน 
	ใช้หลักการคิดคำนวณตามแบบแผนหรือหลักการคิดคำนวณที่คุ้นชินในการแก้ปัญหาคณิตศาสตร์ได้ นักเรียนมีความสามารถในการตีความบริบทที่มีความซับซ้อน แต่มีข้อจำกัดในด้านการแปลความโจทย์ และแปลความจากรูปแบบหนึ่งไปยังรูปแบบหนึ่ง 
	เช่น การแปลความจากรูปภาพเป็นประโยคสัญลักษณ์ทางคณิตศาสตร์ นักเรียนเริ่มมีการใช้เหตุผล ความคิดสร้างสรรค์ รวมถึงจินตนาการเขียนเป็นภาพจำลองสถานการณ์เพื่อช่วยในการแก้ปัญหาคณิตศาสตร์ และนำความรู้ที่มีไปประยุกต์เพื่อแก้ปัญหาคณิตศาสตร์ได้ 
	นักเรียนควรฝึกฝนการวิเคราะห์ปัญหาเชิงลึก และการวางแผนในการแก้ปัญหาอย่างรอบคอบและมีประสิทธิภาพ`,

	"M4": `นักเรียนมีความรู้และความเข้าใจเนื้อหาคณิตศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับดี นักเรียนสามารถคิดวิเคราะห์ และระบุความสัมพันธ์ของปัญหาคณิตศาสตร์ สิ่งที่โจทย์ต้องการทราบ และสิ่งที่โจทย์กำหนดได้อย่างถูกต้องและสมเหตุสมผล 
	รวมถึงสามารถเลือกกระบวนการคิดคำนวณที่มีความหลากหลายได้ นอกจากนี้ นักเรียนสามารถตีความบริบทที่มีความซับซ้อนได้ แต่มีข้อจำกัดในด้านการเชื่อมโยงองค์ประกอบภาพรวม ทั้งนี้ นักเรียนมีการใช้เหตุผล ความคิดสร้างสรรค์ 
	รวมถึงจินตนาการเพื่อเขียนเป็นภาพจำลองสถานการณ์ ซึ่งจัดเป็นหนึ่งในเครื่องมือที่ช่วยในการแก้ปัญหาคณิตศาสตร์ และนำความรู้ที่มีไปประยุกต์เพื่อแก้ปัญหาได้ดี`,

	"M5": `นักเรียนมีความรู้และความเข้าใจเนื้อหาคณิตศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับดีเยี่ยม นักเรียนสามารถคิดวิเคราะห์ และระบุความสัมพันธ์ของปัญหาคณิตศาสตร์ สิ่งที่โจทย์ต้องการทราบ และสิ่งที่โจทย์กำหนดได้อย่างถูกต้องครบถ้วน 
	รวมถึงสามารถเปรียบเทียบ และเลือกกระบวนการคิดคำนวณที่มีความหลากหลายได้อย่างเชี่ยวชาญ นักเรียนมีการใช้เหตุผล ความคิดสร้างสรรค์ รวมถึงจินตนาการเพื่อเขียนเป็นภาพจำลองสถานการณ์ ซึ่งจัดเป็นหนึ่งในเครื่องมือที่ช่วยในการแก้ปัญหา นอกจากนี้ 
	นักเรียนสามารถตีความบริบทที่มีความซับซ้อน การคำนวณเชิงซ้อนหลายขั้นตอน และนำความรู้ที่มีไปประยุกต์เพื่อแก้ปัญหาได้อย่างมีประสิทธิภาพ รวมถึงสามารถบูรณาการกับความรู้แขนงอื่น ๆ เพื่อใช้ในการแก้โจทย์ปัญหาคณิตศาสตร์ในรูปแบบที่แตกต่างกันตามสถานการณ์`,
}

var sciDesp = map[string]string{
	"S1": `นักเรียนมีความรู้และความเข้าใจเนื้อหาวิทยาศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับที่ควรพัฒนา มีทักษะการตีความและสรุปความโดยอาศัยความรู้วิทยาศาสตร์ตามหลักสูตรและความรู้วิทยาศาสตร์จากชีวิตประจำวันมาปรับใช้กับปัญหาได้ในขั้นต้น 
	มีความถูกต้องครบถ้วนเพียงบางส่วน นักเรียนสามารถใช้เพียงความรู้เบื้องต้นในชีวิตประจำวันและกระบวนการตามแบบแผน หรือที่คุ้นชินมาใช้ในการแก้ปัญหา นักเรียนควรพัฒนาทักษะการคิด การวิเคราะห์ความสัมพันธ์ 
	และการเปรียบเทียบข้อมูลโดยใช้กระบวนการทางวิทยาศาสตร์ร่วมกับเหตุผลในการตัดสินใจ รวมถึงเสริมทักษะการตีความ และสรุปความข้อมูลที่มีความซับซ้อน ในขณะเดียวกัน นักเรียนควรเรียนรู้การวางแผนการแก้ปัญหาอย่างเป็นระบบเพื่อนำไปประยุกต์ใช้กับข้อมูล 
	หรือปัญหาที่มีความซับซ้อนให้สามารถไขข้อเท็จจริงได้อย่างถูกต้อง แม่นยำ และมีประสิทธิภาพ`,

	"S2": `นักเรียนมีความรู้และความเข้าใจเนื้อหาวิทยาศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับพอใช้ มีทักษะการตีความและสรุปความโดยอาศัยความรู้วิทยาศาสตร์ตามหลักสูตรควบคู่กับความรู้วิทยาศาสตร์จากชีวิตประจำวันมาปรับใช้กับปัญหาได้ในระดับเบื้องต้น 
	มีทักษะการคิด และวิเคราะห์ปรากฏการณ์หรือสถานการณ์ทางวิทยาศาสตร์ที่มีความคุ้นชินหรือพบในชีวิตประจำวัน นักเรียนควรพัฒนาทักษะการวิเคราะห์ความสัมพันธ์ของปัญหาหรือข้อมูลที่มีความซับซ้อน รวมถึงการใช้เหตุผลประกอบการตัดสินใจเพื่อไขข้อเท็จจริงได้อย่างถูกต้องและแม่นยำ 
	นักเรียนควรเสริมทักษะการนำความรู้พื้นฐานไปประยุกต์ใช้กับศาสตร์แขนงอื่นเพื่อต่อยอดแนวการคิดแก้ปัญหา`,

	"S3": `นักเรียนมีความรู้และความเข้าใจเนื้อหาวิทยาศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับปานกลาง มีทักษะการตีความและสรุปความโดยอาศัยความรู้วิทยาศาสตร์ทั้งตามหลักสูตรและความรู้วิทยาศาสตร์จากชีวิตประจำวันมาปรับใช้กับปัญหาได้ในระดับมาตรฐาน 
	มีทักษะการคิด และวิเคราะห์ปรากฏการณ์หรือสถานการณ์ทางวิทยาศาสตร์ที่มีความคุ้นชินหรือพบในชีวิตประจำวัน และเริ่มมีความสามารถในการวิเคราะห์ข้อมูลที่ซับซ้อนได้ ในขณะเดียวกัน นักเรียนเริ่มมีการเปรียบเทียบข้อมูลโดยใช้กระบวนการทางวิทยาศาสตร์
	ร่วมกับเหตุผลในการตัดสินใจได้ในเบื้องต้น รวมถึงสามารถนำความรู้ไปประยุกต์ใช้เพื่อไขข้อเท็จจริงได้อย่างถูกต้อง และสามารถบูรณาการกับความรู้ศาสตร์แขนงอื่นได้ในระดับปานกลาง`,

	"S4": `นักเรียนมีความรู้และความเข้าใจเนื้อหาวิทยาศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับดี มีทักษะการตีความและสรุปความโดยอาศัยความรู้วิทยาศาสตร์ทั้งตามหลักสูตรและความรู้วิทยาศาสตร์จากชีวิตประจำวันมาปรับใช้กับปัญหาได้อย่างสมเหตุสมผล 
	มีทักษะการคิด วิเคราะห์ และสังเคราะห์ปรากฏการณ์หรือสถานการณ์ทางวิทยาศาสตร์ทั้งรูปแบบทั่วไป และรูปแบบที่มีความซับซ้อนได้อย่างถูกต้อง แต่อาจมีข้อจำกัดด้านความแม่นยำ ในขณะเดียวกัน นักเรียนสามารถแยกแยะ เปรียบเทียบข้อมูล 
	และใช้กระบวนการทางวิทยาศาสตร์ร่วมกับเหตุผลในการตัดสินใจ และนำความรู้ไปประยุกต์ใช้เพื่อไขข้อเท็จจริงได้อย่างถูกต้อง สมเหตุสมผล และสามารถบูรณาการกับความรู้ศาสตร์แขนงอื่นได้ในระดับดี`,

	"S5": `นักเรียนมีความรู้และความเข้าใจเนื้อหาวิทยาศาสตร์ในขอบเขตช่วงชั้นที่กำลังศึกษาระดับดีเยี่ยม มีทักษะการตีความและสรุปความโดยอาศัยความรู้วิทยาศาสตร์ทั้งตามหลักสูตรและความรู้วิทยาศาสตร์จากชีวิตประจำวันมาปรับใช้กับปัญหาได้อย่างชำนาญและสมเหตุสมผล 
	มีทักษะการคิด วิเคราะห์ และสังเคราะห์ปรากฏการณ์หรือสถานการณ์ทางวิทยาศาสตร์ทั้งรูปแบบทั่วไป รูปแบบที่มีความหลากหลาย และรูปแบบที่มีความซับซ้อนได้อย่างชำนาญและแม่นยำ ในขณะเดียวกัน นักเรียนสามารถแยกแยะ เปรียบเทียบข้อมูล 
	และใช้กระบวนการทางวิทยาศาสตร์ร่วมกับเหตุผลในการตัดสินใจ และนำความรู้ไปประยุกต์ใช้เพื่อไขข้อเท็จจริงได้อย่างถูกต้อง ครบถ้วน แม่นยำ สมเหตุสมผล และสามารถบูรณาการกับความรู้ศาสตร์แขนงอื่นได้ในระดับดีเยี่ยม`,
}

var engDesp = map[string]string{
	"E1": `นักเรียนมีความรู้และความเข้าใจเนื้อหาภาษาอังกฤษในขอบเขตช่วงชั้นที่กำลังศึกษาระดับที่ควรพัฒนา นักเรียนรู้และเข้าใจคำศัพท์ โครงสร้างไวยากรณ์ บทสนทนา และบทความที่มีการใช้สำนวนที่ไม่ซับซ้อน 
	สามารถเข้าใจได้จากความคุ้นชิน หรือเคยพบในชีวิตประจำวัน โดยมีความถูกต้องครบถ้วนเพียงบางส่วน นอกจากนี้ นักเรียนสามารถตีความบริบทที่ไม่มีความซับซ้อน โดยอาศัยความรู้ภาษาอังกฤษตามหลักสูตรควบคู่กับความรู้ภาษาอังกฤษจากชีวิตประจำวันได้เบื้องต้น 
	นักเรียนควรพัฒนาทักษะการวิเคราะห์ข้อสอบที่บริบทมีความหมายโดยตรงและความหมายโดยนัยเพื่อระบุความสอดคล้องของโจทย์กับความรู้ภาษาอังกฤษ รวมถึงการเพิ่มพูนความรู้ขั้นพื้นฐาน อาทิ สำนวน โครงสร้างไวยากรณ์ การอ่าน 
	และคำศัพท์เพื่อให้สามารถตัดสินใจเลือกคำตอบได้อย่างถูกต้อง และสมเหตุสมผล`,

	"E2": `นักเรียนมีความรู้และความเข้าใจเนื้อหาภาษาอังกฤษในขอบเขตช่วงชั้นที่กำลังศึกษาระดับพอใช้ นักเรียนรู้และเข้าใจคำศัพท์ รูปแบบประโยคบทสนทนา การโต้ตอบ หรือการจำลองสถานการณ์ที่มีการใช้สำนวนที่คุ้นชิน 
	หรือเคยพบในชีวิตประจำวัน นักเรียนสามารถคิด วิเคราะห์ ตีความ และสรุปใจความบริบทของบทความได้ถูกต้องเพียงบางส่วน โดยมีข้อจำกัดด้านความเข้าใจบริบทภาพรวม ในขณะเดียวกัน นักเรียนรู้และเข้าใจภาษาอังกฤษพื้นฐานในเชิงโครงสร้างประโยคและหลักไวยากรณ์ได้ในระดับพอใช้ 
	นักเรียนควรพัฒนาทักษะการคิดวิเคราะห์ และเพิ่มพูนความรู้ขั้นพื้นฐานทั้งสำนวน โครงสร้างไวยากรณ์ การอ่าน และคำศัพท์เพื่อให้สามารถตัดสินใจเลือกคำตอบที่ถูกต้องได้สมเหตุสมผล`,

	"E3": `นักเรียนมีความรู้และความเข้าใจเนื้อหาภาษาอังกฤษในขอบเขตช่วงชั้นที่กำลังศึกษาระดับปานกลาง นักเรียนรู้และเข้าใจคำศัพท์ รูปแบบประโยคบทสนทนา การโต้ตอบ หรือการจำลองสถานการณ์ที่มีการใช้สำนวนได้ในระดับพื้นฐาน 
	นักเรียนสามารถเริ่มคิด วิเคราะห์ ตีความ และสรุปใจความบริบทของบทความที่มีความหมายโดยนัยได้ นอกจากนี้ นักเรียนสามารถระบุประเภทและหน้าที่ของคำ โครงสร้างประโยค หลักไวยากรณ์ และองค์ประกอบของคำถามได้ในระดับพื้นฐาน 
	ในขณะเดียวกัน นักเรียนมีทักษะในการระบุความสัมพันธ์ของโจทย์กับความรู้พื้นฐานได้อย่างเหมาะสม รวมถึงสามารถนำความรู้เชิงทฤษฎีควบคู่กับประสบการณ์มาประยุกต์ใช้ในการทำข้อสอบได้ในระดับปานกลาง`,

	"E4": `นักเรียนมีความรู้และความเข้าใจเนื้อหาภาษาอังกฤษในขอบเขตช่วงชั้นที่กำลังศึกษาระดับดี นักเรียนรู้และเข้าใจคำศัพท์ รูปแบบประโยคบทสนทนา การโต้ตอบ หรือการจำลองสถานการณ์ที่มีการใช้สำนวน 
	นักเรียนสามารถคิด วิเคราะห์ ตีความ และสรุปใจความบริบทของบทความที่มีความหมายโดยตรงและความหมายโดยนัยได้ในระดับดี นอกจากนี้ นักเรียนสามารถระบุประเภทและหน้าที่ของคำ โครงสร้างประโยค หลักไวยากรณ์ และองค์ประกอบของคำถามได้อย่างถูกต้อง 
	และแม่นยำ ในขณะเดียวกัน นักเรียนมีทักษะในการระบุความสัมพันธ์ของโจทย์กับความรู้พื้นฐานได้อย่างสมเหตุสมผล รวมถึงสามารถนำความรู้เชิงทฤษฎีควบคู่กับประสบการณ์มาประยุกต์ใช้ในการทำข้อสอบได้ในระดับดี`,

	"E5": `นักเรียนมีความรู้และความเข้าใจเนื้อหาภาษาอังกฤษในขอบเขตช่วงชั้นที่กำลังศึกษาระดับดีเยี่ยม นักเรียนรู้และเข้าใจคำศัพท์ รูปแบบประโยคบทสนทนา การโต้ตอบ หรือการจำลองสถานการณ์ที่มีการใช้สำนวน นักเรียนสามารถคิด วิเคราะห์ ตีความ 
	และสรุปใจความบริบทของบทความที่มีความหมายโดยตรงและความหมายโดยนัยได้ในระดับดีเยี่ยม นอกจากนี้ นักเรียนสามารถระบุประเภทและหน้าที่ของคำ โครงสร้างประโยค หลักไวยากรณ์ และองค์ประกอบของคำถามได้อย่างคล่องแคล่ว ถูกต้อง และแม่นยำ 
	ในขณะเดียวกัน นักเรียนสามารถระบุความสัมพันธ์ของโจทย์กับความรู้พื้นฐานได้อย่างสมเหตุสมผล รวมถึงสามารถนำความรู้เชิงทฤษฎีควบคู่กับประสบการณ์มาประยุกต์ใช้ในการทำข้อสอบและศาสตร์แขนงอื่นได้อย่างเชี่ยวชาญ`,
}

func prefixFilter(prefix string, level int64) (string, string) {
	if prefix == "M" {
		return fmt.Sprintf("%s%d", prefix, level), mathDesp[fmt.Sprintf("%s%d", prefix, level)]
	}
	if prefix == "S" {
		return fmt.Sprintf("%s%d", prefix, level), sciDesp[fmt.Sprintf("%s%d", prefix, level)]
	}
	if prefix == "E" {
		return fmt.Sprintf("%s%d", prefix, level), engDesp[fmt.Sprintf("%s%d", prefix, level)]
	}
	return "-", "-"
}
func getClassification(prefix string, score float64) (classification string, desc string) {

	if score >= 0 && score <= 25.0 {
		return prefixFilter(prefix, 1)
	}
	if score >= 25.01 && score <= 49.99 {
		return prefixFilter(prefix, 2)
	}
	if score >= 50.00 && score <= 69.99 {
		return prefixFilter(prefix, 3)
	}
	if score >= 70.00 && score <= 84.99 {
		return prefixFilter(prefix, 4)
	}
	if score >= 85.00 && score <= 100.0 {
		return prefixFilter(prefix, 5)
	}

	return "-", "-"
}
func GetMathAnalytic(c *fiber.Ctx) error {

	payload := &models.GetMathAnalyticRequest{}
	err := c.BodyParser(payload)
	if err != nil {
		return c.SendStatus(400)
	}

	err = utils.Validator().Struct(payload)
	if err != nil {
		return c.SendStatus(400)
	}

	result := &models.MathAnalytic{}
	result.Classification, result.Desc = getClassification("M", float64(payload.ScorePercentage))

	// CAL Part

	if payload.CalPartScore >= 0 && payload.CalPartScore <= 5.65 {
		result.Parts.Calculation = "M1"
	}
	if payload.CalPartScore >= 5.66 && payload.CalPartScore <= 11.28 {
		result.Parts.Calculation = "M2"
	}
	if payload.CalPartScore >= 11.29 && payload.CalPartScore <= 15.80 {
		result.Parts.Calculation = "M3"
	}
	if payload.CalPartScore >= 15.81 && payload.CalPartScore <= 19.19 {
		result.Parts.Calculation = "M4"
	}
	if payload.CalPartScore >= 19.20 && payload.CalPartScore <= 22.60 {
		result.Parts.Calculation = "M5"
	}

	// Problem Solving

	if payload.ProblemPartScore >= 0 && payload.ProblemPartScore <= 13.16 {
		result.Parts.ProblemSolution = "M1"
	}
	if payload.ProblemPartScore >= 13.17 && payload.ProblemPartScore <= 26.27 {
		result.Parts.ProblemSolution = "M2"
	}
	if payload.ProblemPartScore >= 26.28 && payload.ProblemPartScore <= 36.80 {
		result.Parts.ProblemSolution = "M3"
	}
	if payload.ProblemPartScore >= 36.81 && payload.ProblemPartScore <= 44.70 {
		result.Parts.ProblemSolution = "M4"
	}
	if payload.ProblemPartScore >= 44.71 && payload.ProblemPartScore <= 52.65 {
		result.Parts.ProblemSolution = "M5"
	}

	// Applied Part

	if payload.AppliedPartScore >= 0 && payload.AppliedPartScore <= 6.19 {
		result.Parts.Appliation = "M1"
	}
	if payload.AppliedPartScore >= 6.20 && payload.AppliedPartScore <= 12.35 {
		result.Parts.Appliation = "M2"
	}
	if payload.AppliedPartScore >= 12.36 && payload.AppliedPartScore <= 17.30 {
		result.Parts.Appliation = "M3"
	}
	if payload.AppliedPartScore >= 17.31 && payload.AppliedPartScore <= 21.02 {
		result.Parts.Appliation = "M4"
	}
	if payload.AppliedPartScore >= 21.03 && payload.AppliedPartScore <= 24.75 {
		result.Parts.Appliation = "M5"
	}

	return utils.SendSuccess(c, result)

}

func GetSciAnalytic(c *fiber.Ctx) error {
	payload := &models.GetSciAnalyticRequest{}
	err := c.BodyParser(payload)
	if err != nil {
		return c.SendStatus(400)
	}
	err = utils.Validator().Struct(payload)
	if err != nil {
		return c.SendStatus(400)
	}

	result := &models.SciAnalytic{}

	result.Classification, result.Desc = getClassification("S", float64(payload.ScorePercentage))
	// Lesson Part

	if payload.LessonPartScore >= 0 && payload.LessonPartScore <= 20.13 {
		result.Parts.Lesson = "S1"
	}
	if payload.LessonPartScore >= 20.14 && payload.LessonPartScore <= 40.17 {
		result.Parts.Lesson = "S2"
	}
	if payload.LessonPartScore >= 40.18 && payload.LessonPartScore <= 56.27 {
		result.Parts.Lesson = "S3"
	}
	if payload.LessonPartScore >= 56.28 && payload.LessonPartScore <= 68.34 {
		result.Parts.Lesson = "S4"
	}
	if payload.LessonPartScore >= 68.35 && payload.LessonPartScore <= 80.5 {
		result.Parts.Lesson = "S5"
	}

	// Applied Part

	if payload.AppliedPartScore >= 0 && payload.AppliedPartScore <= 4.88 {
		result.Parts.Appliation = "S1"
	}
	if payload.AppliedPartScore >= 4.89 && payload.AppliedPartScore <= 9.73 {
		result.Parts.Appliation = "S2"
	}
	if payload.AppliedPartScore >= 9.74 && payload.AppliedPartScore <= 13.63 {
		result.Parts.Appliation = "S3"
	}
	if payload.AppliedPartScore >= 13.64 && payload.AppliedPartScore <= 16.56 {
		result.Parts.Appliation = "S4"
	}
	if payload.AppliedPartScore >= 16.57 && payload.AppliedPartScore <= 1950 {
		result.Parts.Appliation = "S5"
	}

	return utils.SendSuccess(c, result)
}

func GetEngAnalytic(c *fiber.Ctx) error {
	payload := &models.GetEngAnalyticRequest{}
	err := c.BodyParser(payload)
	if err != nil {
		return c.SendStatus(400)
	}
	err = utils.Validator().Struct(payload)
	if err != nil {
		return c.SendStatus(400)
	}

	result := &models.EngAnalytic{}

	result.Classification, result.Desc = getClassification("E", float64(payload.ScorePercentage))

	// Expression Part

	if payload.ExpressionPartScore >= 0 && payload.ExpressionPartScore <= 4.00 {
		result.Parts.Expression = "E1"
	}
	if payload.ExpressionPartScore >= 4.10 && payload.ExpressionPartScore <= 7.98 {
		result.Parts.Expression = "E2"
	}
	if payload.ExpressionPartScore >= 7.99 && payload.ExpressionPartScore <= 11.18 {
		result.Parts.Expression = "E3"
	}
	if payload.ExpressionPartScore >= 11.19 && payload.ExpressionPartScore <= 13.58 {
		result.Parts.Expression = "E4"
	}
	if payload.ExpressionPartScore >= 13.59 && payload.ExpressionPartScore <= 16.00 {
		result.Parts.Expression = "E5"
	}
	// Reading Part

	if payload.ReadingPartScore >= 0 && payload.ReadingPartScore <= 9.00 {
		result.Parts.Reading = "E1"
	}
	if payload.ReadingPartScore >= 9.01 && payload.ReadingPartScore <= 17.96 {
		result.Parts.Reading = "E2"
	}
	if payload.ReadingPartScore >= 17.97 && payload.ReadingPartScore <= 25.16 {
		result.Parts.Reading = "E3"
	}
	if payload.ReadingPartScore >= 25.17 && payload.ReadingPartScore <= 30.56 {
		result.Parts.Reading = "E4"
	}
	if payload.ReadingPartScore >= 30.57 && payload.ReadingPartScore <= 36.00 {
		result.Parts.Reading = "E5"
	}

	// Structure

	if payload.StructPartScore >= 0 && payload.StructPartScore <= 8.25 {
		result.Parts.Structure = "E1"
	}
	if payload.StructPartScore >= 8.26 && payload.StructPartScore <= 16.47 {
		result.Parts.Structure = "E2"
	}
	if payload.StructPartScore >= 16.48 && payload.StructPartScore <= 23.07 {
		result.Parts.Structure = "E3"
	}
	if payload.StructPartScore >= 23.08 && payload.StructPartScore <= 28.02 {
		result.Parts.Structure = "E4"
	}
	if payload.StructPartScore >= 28.03 && payload.StructPartScore <= 33.0 {
		result.Parts.Structure = "E5"
	}

	// Vocabuary

	if payload.VocabularyPartScore >= 0 && payload.VocabularyPartScore <= 3.75 {
		result.Parts.Vocabulary = "E1"
	}
	if payload.VocabularyPartScore >= 3.76 && payload.VocabularyPartScore <= 7.49 {
		result.Parts.Vocabulary = "E2"
	}
	if payload.VocabularyPartScore >= 7.50 && payload.VocabularyPartScore <= 10.49 {
		result.Parts.Vocabulary = "E3"
	}
	if payload.VocabularyPartScore >= 10.50 && payload.VocabularyPartScore <= 12.74 {
		result.Parts.Vocabulary = "E4"
	}
	if payload.VocabularyPartScore >= 12.75 && payload.VocabularyPartScore <= 15.00 {
		result.Parts.Vocabulary = "E5"
	}

	return utils.SendSuccess(c, result)

}

type getIaarDataRequest struct {
	Year    string `json:"year" validate:"required"`
	HashCid string `json:"hash_cid" validate:"required"`
	Subject string `json:"subject" validate:"required,eq=ENG|eq=MATH|eq=SCI"`
}

func GetIaarData(c *fiber.Ctx) error {

	req := &getIaarDataRequest{}
	err := c.BodyParser(req)
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 4000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Failed to parse request",
				ErrorMessage: err.Error(),
			},
		})
	}

	err = utils.Validator().Struct(req)
	if err != nil {
		return utils.SendCommonError(c, models.CommonError{
			Code: 4000,
			ErrorData: models.ApiError{
				ErrorTitle:   "Failed to validate request",
				ErrorMessage: err.Error(),
			},
		})
	}

	// Get IAAR Data
	db := database.DB.Db

	if req.Subject == "ENG" {
		var iaar models.EngIaar
		sql := `
		select
			es.hash_cid,
			c."name",
			c.level_range ,
			c.school ,
			c.province ,
			c.exam_type ,
			c.region,
			es.score_pt_expression ,
			es.score_pt_reading ,
			es.score_pt_structure ,
			es.score_pt_vocabulary ,
			es.total_score ,
			es.province_rank ,
			es.region_rank
		from
			eng_scores es
		inner join competitors c 
		on
			c.cid = es.hash_cid
		where
			c.cid = ?;
	`

		err = db.Raw(sql, req.HashCid).Scan(&iaar).Error
		if err != nil {
			slog.Error("Failed to get iaar data", "err", err.Error())
		}

		if iaar.HashCid == "" {
			return utils.SendCommonError(c, models.CommonError{
				Code: 4000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Not found",
					ErrorMessage: "Not found",
				},
			})
		}

		slr, err := utils.GetShortLevelRange(iaar.LevelRange)
		if err != nil {
			slog.Error("Failed to get short level range", "err", err.Error())
		}
		iaar.ShortLevelRange = slr

		avgScore := &models.AvgScoreBySubject{}
		err = db.Where("level_range = ? AND subject = ? AND year = ?", slr, req.Subject, req.Year).Find(avgScore).Error
		if err != nil {
			slog.Error("Failed to get avg score", "err", err.Error())
			return utils.SendCommonError(c, models.CommonError{
				Code: 5000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Failed to get avg score",
					ErrorMessage: err.Error(),
				},
			})
		}
		iaar.RegionMaxScore = avgScore.MaxScore
		iaar.RegionAvgScore = avgScore.AvgScore

		numOfCompByProvince := &models.NumberOfCompetitorByProvince{}
		err = db.Where("level_range = ? AND subject = ? AND year = ? AND province = ?", slr, req.Subject, req.Year, iaar.Province).Find(numOfCompByProvince).Error
		if err != nil {
			slog.Error("Failed to get num of competitor by province", "err", err.Error())
			return utils.SendCommonError(c, models.CommonError{
				Code: 5000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Failed to get num of competitor by province",
					ErrorMessage: err.Error(),
				},
			})
		}
		iaar.ProvinceRank = fmt.Sprintf("%v/%v", iaar.ProvinceRank, numOfCompByProvince.NumberOfCompetitor)

		numOfCompByRegion := &models.NumberOfCompetitorByRegion{}
		err = db.Where("level_range = ? AND subject = ? AND year = ? AND region = ?", slr, req.Subject, req.Year, iaar.Region).Find(numOfCompByRegion).Error
		if err != nil {
			slog.Error("Failed to get num of competitor by region", "err", err.Error())
			return utils.SendCommonError(c, models.CommonError{
				Code: 5000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Failed to get num of competitor by region",
					ErrorMessage: err.Error(),
				},
			})
		}
		iaar.RegionRank = fmt.Sprintf("%v/%v", iaar.RegionRank, numOfCompByRegion.NumberOfCompetitor)

		iaar.AnalyticData = getEngAnalytic(&models.EngScorePerPart{
			ScorePtExpression: iaar.ScorePtExpression,
			ScorePtReading:    iaar.ScorePtReading,
			ScorePtStructure:  iaar.ScorePtStructure,
			ScorePtVocabulary: iaar.ScorePtVocabulary,
		}, iaar.TotalScore)

		iaar.PrizeTypeTH, iaar.PrizeTypeEN = getPrizeType(iaar.TotalScore)

		return utils.SendSuccess(c, iaar)
	}

	if req.Subject == "MATH" {
		var iaar models.MathIaar
		sql := `
		select
			ms.hash_cid,
			c."name",
			c.level_range ,
			c.school ,
			c.province ,
			c.exam_type ,
			c.region,
			ms.score_pt_calculate ,
			ms.score_pt_problem_math ,
			ms.score_pt_applied_math ,
			ms.total_score ,
			ms.province_rank ,
			ms.region_rank
		from
			math_scores ms
		inner join competitors c 
		on
			c.cid = ms.hash_cid
		where
			c.cid = ?;
	`

		err = db.Raw(sql, req.HashCid).Scan(&iaar).Error
		if err != nil {
			slog.Error("Failed to get iaar data", "err", err.Error())
		}

		if iaar.HashCid == "" {
			return utils.SendCommonError(c, models.CommonError{
				Code: 4000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Not found",
					ErrorMessage: "Not found",
				},
			})
		}

		slr, err := utils.GetShortLevelRange(iaar.LevelRange)
		if err != nil {
			slog.Error("Failed to get short level range", "err", err.Error())
		}
		iaar.ShortLevelRange = slr

		avgScore := &models.AvgScoreBySubject{}
		err = db.Where("level_range = ? AND subject = ? AND year = ?", slr, req.Subject, req.Year).Find(avgScore).Error
		if err != nil {
			slog.Error("Failed to get avg score", "err", err.Error())
			return utils.SendCommonError(c, models.CommonError{
				Code: 5000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Failed to get avg score",
					ErrorMessage: err.Error(),
				},
			})
		}
		iaar.RegionMaxScore = avgScore.MaxScore
		iaar.RegionAvgScore = avgScore.AvgScore

		numOfCompByProvince := &models.NumberOfCompetitorByProvince{}
		err = db.Where("level_range = ? AND subject = ? AND year = ? AND province = ?", slr, req.Subject, req.Year, iaar.Province).Find(numOfCompByProvince).Error
		if err != nil {
			slog.Error("Failed to get num of competitor by province", "err", err.Error())
			return utils.SendCommonError(c, models.CommonError{
				Code: 5000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Failed to get num of competitor by province",
					ErrorMessage: err.Error(),
				},
			})
		}
		iaar.ProvinceRank = fmt.Sprintf("%v/%v", iaar.ProvinceRank, numOfCompByProvince.NumberOfCompetitor)

		numOfCompByRegion := &models.NumberOfCompetitorByRegion{}
		err = db.Where("level_range = ? AND subject = ? AND year = ? AND region = ?", slr, req.Subject, req.Year, iaar.Region).Find(numOfCompByRegion).Error
		if err != nil {
			slog.Error("Failed to get num of competitor by region", "err", err.Error())
			return utils.SendCommonError(c, models.CommonError{
				Code: 5000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Failed to get num of competitor by region",
					ErrorMessage: err.Error(),
				},
			})
		}
		iaar.RegionRank = fmt.Sprintf("%v/%v", iaar.RegionRank, numOfCompByRegion.NumberOfCompetitor)

		iaar.AnalyticData = getMathAnalytic(&models.MathScorePerPart{
			ScorePtCalculate:   iaar.ScorePtCalculate,
			ScorePtProblemMath: iaar.ScorePtProblemMath,
			ScorePtAppliedMath: iaar.ScorePtAppliedMath,
		}, iaar.TotalScore)

		iaar.PrizeTypeTH, iaar.PrizeTypeEN = getPrizeType(iaar.TotalScore)
		iaar.Subject = "MATH"

		return utils.SendSuccess(c, iaar)
	}
	if req.Subject == "SCI" {
		var iaar models.SciIaar
		sql := `
		select
			ss.hash_cid,
			c."name",
			c.level_range ,
			c.school ,
			c.province ,
			c.exam_type ,
			c.region,
			ss.score_pt_lesson_sci ,
			ss.score_pt_applied_sci ,
			ss.total_score ,
			ss.province_rank ,
			ss.region_rank
		from
			sci_scores ss
		inner join competitors c
		on
			c.cid = ss.hash_cid
		where
			c.cid = ?;
	`

		err = db.Raw(sql, req.HashCid).Scan(&iaar).Error
		if err != nil {
			slog.Error("Failed to get iaar data", "err", err.Error())
		}

		if iaar.HashCid == "" {
			return utils.SendCommonError(c, models.CommonError{
				Code: 4000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Not found",
					ErrorMessage: "Not found",
				},
			})
		}

		slr, err := utils.GetShortLevelRange(iaar.LevelRange)
		if err != nil {
			slog.Error("Failed to get short level range", "err", err.Error())
		}
		iaar.ShortLevelRange = slr

		avgScore := &models.AvgScoreBySubject{}
		err = db.Where("level_range = ? AND subject = ? AND year = ?", slr, req.Subject, req.Year).Find(avgScore).Error
		if err != nil {
			slog.Error("Failed to get avg score", "err", err.Error())
			return utils.SendCommonError(c, models.CommonError{
				Code: 5000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Failed to get avg score",
					ErrorMessage: err.Error(),
				},
			})
		}
		iaar.RegionMaxScore = avgScore.MaxScore
		iaar.RegionAvgScore = avgScore.AvgScore

		numOfCompByProvince := &models.NumberOfCompetitorByProvince{}
		err = db.Where("level_range = ? AND subject = ? AND year = ? AND province = ?", slr, req.Subject, req.Year, iaar.Province).Find(numOfCompByProvince).Error
		if err != nil {
			slog.Error("Failed to get num of competitor by province", "err", err.Error())
			return utils.SendCommonError(c, models.CommonError{
				Code: 5000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Failed to get num of competitor by province",
					ErrorMessage: err.Error(),
				},
			})
		}
		iaar.ProvinceRank = fmt.Sprintf("%v/%v", iaar.ProvinceRank, numOfCompByProvince.NumberOfCompetitor)

		numOfCompByRegion := &models.NumberOfCompetitorByRegion{}
		err = db.Where("level_range = ? AND subject = ? AND year = ? AND region = ?", slr, req.Subject, req.Year, iaar.Region).Find(numOfCompByRegion).Error
		if err != nil {
			slog.Error("Failed to get num of competitor by region", "err", err.Error())
			return utils.SendCommonError(c, models.CommonError{
				Code: 5000,
				ErrorData: models.ApiError{
					ErrorTitle:   "Failed to get num of competitor by region",
					ErrorMessage: err.Error(),
				},
			})
		}
		iaar.RegionRank = fmt.Sprintf("%v/%v", iaar.RegionRank, numOfCompByRegion.NumberOfCompetitor)

		iaar.AnalyticData = getSciAnalytic(&models.SciScorePerPart{
			ScorePtLessonSci:  iaar.ScorePtLessonSci,
			ScorePtAppliedSci: iaar.ScorePtAppliedSci,
		}, iaar.TotalScore)

		iaar.PrizeTypeTH, iaar.PrizeTypeEN = getPrizeType(iaar.TotalScore)
		iaar.Subject = "SCI"

		return utils.SendSuccess(c, iaar)
	}

	return nil
}

func getEngAnalytic(score *models.EngScorePerPart, totalScore float64) models.EngAnalytic {

	result := models.EngAnalytic{}
	result.Classification, result.Desc = getClassification("E", totalScore)

	// Expression Part
	if score.ScorePtExpression >= 0 && score.ScorePtExpression <= 4.00 {
		result.Parts.Expression = "E1"
	}
	if score.ScorePtExpression >= 4.10 && score.ScorePtExpression <= 7.98 {
		result.Parts.Expression = "E2"
	}
	if score.ScorePtExpression >= 7.99 && score.ScorePtExpression <= 11.18 {
		result.Parts.Expression = "E3"
	}
	if score.ScorePtExpression >= 11.19 && score.ScorePtExpression <= 13.58 {
		result.Parts.Expression = "E4"
	}
	if score.ScorePtExpression >= 13.59 && score.ScorePtExpression <= 16.00 {
		result.Parts.Expression = "E5"
	}
	// Reading Part

	if score.ScorePtReading >= 0 && score.ScorePtReading <= 9.00 {
		result.Parts.Reading = "E1"
	}
	if score.ScorePtReading >= 9.01 && score.ScorePtReading <= 17.96 {
		result.Parts.Reading = "E2"
	}
	if score.ScorePtReading >= 17.97 && score.ScorePtReading <= 25.16 {
		result.Parts.Reading = "E3"
	}
	if score.ScorePtReading >= 25.17 && score.ScorePtReading <= 30.56 {
		result.Parts.Reading = "E4"
	}
	if score.ScorePtReading >= 30.57 && score.ScorePtReading <= 36.00 {
		result.Parts.Reading = "E5"
	}

	// Structure

	if score.ScorePtStructure >= 0 && score.ScorePtStructure <= 8.25 {
		result.Parts.Structure = "E1"
	}
	if score.ScorePtStructure >= 8.26 && score.ScorePtStructure <= 16.47 {
		result.Parts.Structure = "E2"
	}
	if score.ScorePtStructure >= 16.48 && score.ScorePtStructure <= 23.07 {
		result.Parts.Structure = "E3"
	}
	if score.ScorePtStructure >= 23.08 && score.ScorePtStructure <= 28.02 {
		result.Parts.Structure = "E4"
	}
	if score.ScorePtStructure >= 28.03 && score.ScorePtStructure <= 33.0 {
		result.Parts.Structure = "E5"
	}

	// Vocabuary

	if score.ScorePtVocabulary >= 0 && score.ScorePtVocabulary <= 3.75 {
		result.Parts.Vocabulary = "E1"
	}
	if score.ScorePtVocabulary >= 3.76 && score.ScorePtVocabulary <= 7.49 {
		result.Parts.Vocabulary = "E2"
	}
	if score.ScorePtVocabulary >= 7.50 && score.ScorePtVocabulary <= 10.49 {
		result.Parts.Vocabulary = "E3"
	}
	if score.ScorePtVocabulary >= 10.50 && score.ScorePtVocabulary <= 12.74 {
		result.Parts.Vocabulary = "E4"
	}
	if score.ScorePtVocabulary >= 12.75 && score.ScorePtVocabulary <= 15.00 {
		result.Parts.Vocabulary = "E5"
	}

	return result
}

func getMathAnalytic(score *models.MathScorePerPart, totalScore float64) models.MathAnalytic {
	result := models.MathAnalytic{}
	result.Classification, result.Desc = getClassification("M", totalScore)

	// CAL Part
	if score.ScorePtCalculate >= 0 && score.ScorePtCalculate <= 5.65 {
		result.Parts.Calculation = "M1"
	}
	if score.ScorePtCalculate >= 5.66 && score.ScorePtCalculate <= 11.28 {
		result.Parts.Calculation = "M2"
	}
	if score.ScorePtCalculate >= 11.29 && score.ScorePtCalculate <= 15.80 {
		result.Parts.Calculation = "M3"
	}
	if score.ScorePtCalculate >= 15.81 && score.ScorePtCalculate <= 19.19 {
		result.Parts.Calculation = "M4"
	}
	if score.ScorePtCalculate >= 19.20 && score.ScorePtCalculate <= 22.60 {
		result.Parts.Calculation = "M5"
	}

	// Problem Solving
	if score.ScorePtProblemMath >= 0 && score.ScorePtProblemMath <= 13.16 {
		result.Parts.ProblemSolution = "M1"
	}
	if score.ScorePtProblemMath >= 13.17 && score.ScorePtProblemMath <= 26.27 {
		result.Parts.ProblemSolution = "M2"
	}
	if score.ScorePtProblemMath >= 26.28 && score.ScorePtProblemMath <= 36.80 {
		result.Parts.ProblemSolution = "M3"
	}
	if score.ScorePtProblemMath >= 36.81 && score.ScorePtProblemMath <= 44.70 {
		result.Parts.ProblemSolution = "M4"
	}
	if score.ScorePtProblemMath >= 44.71 && score.ScorePtProblemMath <= 52.65 {
		result.Parts.ProblemSolution = "M5"
	}

	// Applied Part

	if score.ScorePtAppliedMath >= 0 && score.ScorePtAppliedMath <= 6.19 {
		result.Parts.Appliation = "M1"
	}
	if score.ScorePtAppliedMath >= 6.20 && score.ScorePtAppliedMath <= 12.35 {
		result.Parts.Appliation = "M2"
	}
	if score.ScorePtAppliedMath >= 12.36 && score.ScorePtAppliedMath <= 17.30 {
		result.Parts.Appliation = "M3"
	}
	if score.ScorePtAppliedMath >= 17.31 && score.ScorePtAppliedMath <= 21.02 {
		result.Parts.Appliation = "M4"
	}
	if score.ScorePtAppliedMath >= 21.03 && score.ScorePtAppliedMath <= 24.75 {
		result.Parts.Appliation = "M5"
	}
	return result
}

func getSciAnalytic(score *models.SciScorePerPart, totalScore float64) models.SciAnalytic {
	result := models.SciAnalytic{}
	result.Classification, result.Desc = getClassification("S", totalScore)

	// Lesson Part
	if score.ScorePtLessonSci >= 0 && score.ScorePtLessonSci <= 20.13 {
		result.Parts.Lesson = "S1"
	}
	if score.ScorePtLessonSci >= 20.14 && score.ScorePtLessonSci <= 40.17 {
		result.Parts.Lesson = "S2"
	}
	if score.ScorePtLessonSci >= 40.18 && score.ScorePtLessonSci <= 56.27 {
		result.Parts.Lesson = "S3"
	}
	if score.ScorePtLessonSci >= 56.28 && score.ScorePtLessonSci <= 68.34 {
		result.Parts.Lesson = "S4"
	}
	if score.ScorePtLessonSci >= 68.35 && score.ScorePtLessonSci <= 80.5 {
		result.Parts.Lesson = "S5"
	}

	// Applied Part
	if score.ScorePtAppliedSci >= 0 && score.ScorePtAppliedSci <= 4.88 {
		result.Parts.Appliation = "S1"
	}
	if score.ScorePtAppliedSci >= 4.89 && score.ScorePtAppliedSci <= 9.73 {
		result.Parts.Appliation = "S2"
	}
	if score.ScorePtAppliedSci >= 9.74 && score.ScorePtAppliedSci <= 13.63 {
		result.Parts.Appliation = "S3"
	}
	if score.ScorePtAppliedSci >= 13.64 && score.ScorePtAppliedSci <= 16.56 {
		result.Parts.Appliation = "S4"
	}
	if score.ScorePtAppliedSci >= 16.57 && score.ScorePtAppliedSci <= 1950 {
		result.Parts.Appliation = "S5"
	}

	return result
}

func getPrizeType(score float64) (string, string) {

	if score >= 0 && score < 50 {
		return "รางวัลชมเชย", "CON"
	}
	if score >= 50 && score < 70 {
		return "รางวัลเหรียญทองแดง", "BRONZE"
	}
	if score >= 70 && score < 85 {
		return "รางวัลเหรียญเงิน", "SILVER"
	}
	if score >= 85 && score <= 100 {
		return "รางวัลเหรียญทอง", "GOLD"
	}

	return "-", "-"
}
