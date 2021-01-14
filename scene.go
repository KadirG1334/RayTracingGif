package main
import ()



// RandomScene returns a 'random' scene
func Scene() *World {
	world := &World{}

	floor := NewSphere(0, -1000, 0, 1000,
		Lambertian{Attenuation: Color{
			R: 0,
			G: 1, //1 e 0.20
			B: 0.20}})

	world.Add(floor)



	startsphere:=NewSphere(0,6,-6,2,
	Metal{Attenuation:Color{
		R:0.984375,
		G:0.828125,
		B:0.25},Fuzz:0.1})
		
		world.Add(startsphere)


	rotatingsphere1:=NewSphere(3,1,-6,1,
	Lambertian{Attenuation:Color{
		R:0,
		G:0.5,
		B:0.5}})
    rotatingsphere2:=NewSphere(-3,1,-6,1,
			Lambertian{Attenuation:Color{
				R:0,
				G:0.5,
				B:0.5}})
	rotatingsphere3:=NewSphere(0,1,-3,1,
					Lambertian{Attenuation:Color{
						R:0,
						G:0.5,
						B:0.5}})
	rotatingsphere4:=NewSphere(0,1,-9,1,
							Lambertian{Attenuation:Color{
								R:0,
								G:0.5,
								B:0.5}})
													
																						

				
								if i==0{
									world.Add(rotatingsphere3)
	
								}else if i==1{
									world.Add(rotatingsphere1)
								}else if i==2{
									world.Add(rotatingsphere4)
								}else if i==3{
									world.Add(rotatingsphere2)
								}


	//lambertian := NewSphere(-4, 1, 0, 1.0, Lambertian{Attenuation: Color{R: 0.4, G: 0.0, B: 0.1}}) 
	//metal := NewSphere(4, 1, 0, 1.0, Metal{Attenuation:Color{R: 0.7, G: 0.6, B: 0.5}, Fuzz: 0.8})
	
	
	return world
}