from gpiozero import MotionSensor
import RPi.GPIO as GPIO

pir = MotionSensor(27)

while True:
	pir.wait_for_motion()
	print("You moved")
	pir.wait_for_no_motion()
