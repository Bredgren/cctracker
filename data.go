package main

type standard struct {
	Sets, Reps int
}

type exInfo struct {
	Name        string
	Units       string
	Best        int
	Performance string
	XRay        string
	Goals       [3]standard
	Tips        string
}

type groupInfo struct {
	Name      string
	Exercises []exInfo
}

var exercises = []groupInfo{
	{
		"Pushups",
		[]exInfo{
			{"Wall Pushups", "", 0,
				`Face a wall. With your feet together, place your palms flat against the wall. This is the start position
(fig. 1). Your arms should be straight and shoulder width apart, with the hands at chest level.
Bend the shoulders and elbows until the forehead gently touches the wall. This is the finish position
(fig. 2). Press back to the start position and repeat.`,
				`Wall pushups are the first step of the ten step series required for complete mastery of the pushup
family of movements. As the first step, this technique represents the easiest version of the pushup.
Every able-bodied person should be able to do this exercise without a problem. Wall pushups are
also the first movement in the therapy sequence of the pushup series. This exercise will be of great
benefit to somebody coming out of an injury or following an operation, who is looking to promote
healing and rebuild their strength slowly. The elbows, wrists and shoulders--most notably
the delicate rotator cuff muscles within the shoulders--are particularly prone to chronic and acute
injury. This exercise gently activates these areas, stimulating them, developing blood flow and
tone. Beginners new to calisthenics must start any training program very gently to develop their
technique and ability. They should start with this exercise.`,
				[3]standard{{1, 10}, {2, 25}, {3, 50}},
				`Every person reading this book should be able to do this exercise, unless they are disabled,
badly injured or ill. If coming back from an injury or operation, this movement is a good "tester,"
allowing the athlete to feel out any weak points during rehabilitation.`},
			{"Incline Pushups", "", 0,
				`To do this exercise you will need to find a secure object or piece of furniture which is about half
your height-it should reach the midpoint of your body (roughly around the level of the hips).
Good options might be desks, tall chairs, work surfaces, kitchen tops, low walls or solid fences. In
most prisons, I've found that the cell basin fits the bill nicely--but make sure it's strong enough to
accommodate the demands of the exercise. With your feet together and your body aligned, lean
over and grasp the object with the arms straight and shoulder width apart. This is the start position
of the exercise and--if the object reaches your midpoint--should put you at about 45 degrees
from the floor (see fig. 3). Bending at the elbows and shoulders, lower yourself until your torso
gently touches the top of the object (fig. 4). Pause briefly before pressing back up to the start position,
and repeat.`,
				`This exercise continues where Step 1 (the wall pushup) leaves off, with the lower pressing angle
meaning that more bodyweight has to be moved by the muscles of the upper body. The incline
pushup is easier than the classic full pushup (Step 5). For most athletes this exercise won't place
very great demands on the muscles, but it will be useful for the beginner to continue their training
gently, or for rehabilitation purposes.`,
				[3]standard{{1, 10}, {2, 20}, {3, 40}},
				`Incline pushups should be done at 45 degree angle. If the beginner standard can't be met at this
angle, use a greater (more upright) angle-simply place the hands on an object higher than the
body's midpoint. Once this is mastered, gradually use lesser angles until 45 degrees becomes easy.
Harder angles can be attempted using progressively lower steps on a set of stairs.`},
			{"Kneeling Pushups", "", 0,
				`Kneel on the floor with your feet together, and your palms flat on the ground in front of you.
The arms should be straight, shoulder width apart, and in line with your chest. Link one ankle
around the other, and keep the hips straight and in alignment with the trunk and head. This is the
start position (fig. 5). Using the knees as a pivot, bend at the shoulders and elbows until your chest
is approximately one fist's width from the floor (fig. 6). Pause and press back to the start position,
then repeat.`,
				`Kneeling pushups are Step 3 in the pushup series. They are an important movement for beginners
to master, because they are the easiest type of pushup that can be performed prone, i.e., flat
on the floor. Because of this, they form an important link between the first two steps, which are
performed standing, and the harder prone techniques later in the series. Women often do kneeling
pushups because they lack the relative upper body strength to perform the full pushup, but this
exercise offers great benefits to guys as well. It's a good starting exercise for somebody who's
maybe overweight or out of shape, and because you can pump up the upper body with relative
ease in this position, kneeling pushups make an excellent warm up exercise you can do before
attempting harder forms of pushup.`,
				[3]standard{{1, 10}, {2, 15}, {3, 30}},
				`If you find it impossible to perform full kneeling pushups, lessen your range of motion. Don't go
all the way down to a fist width from the ground. Use a higher number of reps (about twenty)
over a shorter range of motion you can perform comfortably, then workout by workout (keeping
the reps high) gradually keep adding an inch of depth until the full movement is mastered.`},
			{"Half Pushups", "", 0,
				`From the kneeling position, place your palms on the floor and stretch your legs out behind you.
Your hands should be shoulder width apart, and directly below your upper chest. Your feet and
legs should be kept together. Tighten your supporting muscles, so that your back, hips and legs
stay locked in line. Starting with the arms straight, lower yourself approximately half the length of
your extended arms, or until your elbows form a right angle. An excellent way to establish how
far to descend is to use a standard basketball or soccer ball. Position yourself over the ball at the
top of the movement, so that the ball is directly below your hips. This is the start position (fig. 7).
Bend at the shoulders and elbows until your hips lightly make contact with the ball (fig. 8). On
most people, this will be a good, objective indicator of the correct bottom position. Pause before
pressing forcefully back to the start position.`,
				`The half pushup is an important exercise to master properly. You see a lot of guys doing
pushups incorrectly, with their butt moving up as they bend at the hips. They do this because their
waist and spinal muscles are weak. This exercise trains the waist and spine to keep the hips
locked and aligned.`,
				[3]standard{{1, 8}, {2, 12}, {2, 25}},
				`If you can't do half pushups, shorten your range of motion until you are able to perform the
technique as given above. If you are using a basketball to monitor your form, position yourself so
that it's under your knees rather your hips. If you lower yourself from the arms extended position
until your knees make contact with the ball, this will approximately equal a quarter pushup. Once
you can do more than ten quarter pushups, position the ball a little higher up your thighs each
time you practice until it is under your hips.`},
			{"Full Pushups", "", 0,
				`From the kneeling position, place your palms on the floor and stretch your legs out behind you.
Keep your thighs and feet together, and ensure that the hands are below your upper chest and
shoulder width apart. Straighten the arms. The hips and spine should be in line. This is the start
position (fig 9). Bend at the elbows and shoulders until your breastbone comes to within a fist's
width from the floor. In prison pushup competitions, a "counter" clenches his fist pinky side
down on the floor, and counts out when the athlete's chest touches the knuckle of his thumb. If
you're training alone and you wish to keep to the right depth, place a baseball or tennis ball
directly below your chest (fig. 10). As your chest kisses the ball, pause and push up.`,
				`This technique is the "classic" pushup, the exercise most of us will remember from gym class.
It's a fair guess that if you say the word "pushup" to most people, they will naturally picture the
full pushup. The full pushup is an excellent upper body exercise, working the arms, chest and
shoulder girdle in an efficient manner. It's by no means the hardest form of pushup, however; in
terms of difficulty it only represents Step 5 in a series of ten.`,
				[3]standard{{1, 5}, {2, 10}, {2, 20}},
				`You might be surprised how many people--even big, strong guys--eannot do full pushups
properly. If you are in this group, return to half pushups using a basketball. If you have graduated
from Step 4, you will be able to perform twenty-five reps with the ball under your hips. Gradually
move the ball a few inches forward every workout, or whenever you can-keeping the reps the
same. When you can go from the arms straight position to a bottom position where your jaw
touches the ball, attempt the full version again.`},
			{"Close Pushups", "", 0,
				`Begin this technique in the same top position as full pushups (Step 5), but with the hands touching.
You don't need to overlap the hands or form a "diamond" between the thumbs and index fingers;
it is sufficient for the tips of the index fingers to touch. From the straight arm start position
(fig. 11), lower yourself until your chest gently touches the backs of your hands (fig. 12). Pause
briefly before pressing back to the start position.`,
				`Close pushups are as old as the hills. They're a vitally important exercise in the pushup series,
but are often overlooked in favor of flashier techniques like plyometric (clapping) pushups and
decline pushups. This is a tragedy, because the close pushup is an essential tool to help you in your
journey to mastering the one-arm pushup. Most athletes have trouble with the one-arm pushup
because they find it difficult to press themselves up from the bottom position, when the arm is
bent most acutely. This is because their elbows are weak when bent beyond a right angle. Because
of the placement of the hands during close pushups, the athlete naturally bends his elbows to a
greater degree to reach the bottom position than is the case with full pushups. This increased
elbow flexion trains the triceps and strengthens the tendons of the elbow and wrist. As a result,
athletes who have become comfortable with this movement will find one-arm pushups much
more manageable when the time comes.`,
				[3]standard{{1, 5}, {2, 10}, {2, 20}},
				`If you cannot do close pushups with the hands touching (as described above), simply return to
full pushups, and move the hands an inch or two closer every workout, keeping the reps quite high.`},
			{"Uneven Pushups", "", 0,
				`Get into the classic pushup position; feet together, legs, hips and back aligned, and arms straight
with the palms on the floor beneath your upper chest. With one arm firmly supporting you, place
the other hand on a basketball. Both your hands should be directly below your shoulders for stability.
This is the start position (fig. 13). Once you have found your balance, do your best to evenly
distribute your weight through both hands. At first this will not be easy, but persevere. Bend at the
elbows and shoulders until your chest touches the top of whichever hand is on the ball (fig. 14).
Pause briefly before pressing back up to the start position.`,
				`This is the first of the advanced pushup exercises that will take the athlete from double arm
pushups to the single arm variety. You can use stable objects-like bricks or a cinderblock instead
of a basketball, but the basketball is best. The act of stabilizing the ball brings the seldomused
rotator cuff muscles into play, strengthening them for the more intense exercises to come.
You can use a sturdy soccer ball rather than a basketball, but a classic basketball is king because
its tacky surface helps the palm grip.`,
				[3]standard{{1, 5}, {2, 10}, {2, 20}},
				`Anybody who can do close pushups properly should be ready to attempt this exercise with confidence.
If there are problems at first, they will be due to a lack of coordination rather than insufficient
strength. If you have trouble, try using stable objects rather than an unstable basketball--a
simple house brick is a good alternative. Once you can do twenty reps with one hand on a flat
brick, try two flat bricks, one on top of the other. When you can do twenty reps with three stacked
bricks, attempt the exercise with the basketball again.`},
			{"1/2 One-Arm Pushups", "", 0,
				`Get into the half pushup top position, with a basketball located below your hips, as described
in Step 4. Place one hand on the floor beneath your breastbone with your arm straight and your
other hand in the small of your back. This is the start position (fig. 15). Bend at the shoulder and
elbow until your hips reach the top of the basketball. This is the finish position (fig. 16). Pause and
press back to the start position. If your triceps is weak, you'll have a tendency to twist your torso
as you move. Don't-the whole body should be kept straight. This is true for all pushups.`,
				`Half one-arm pushups are Step 8 of the pushup series of movements. With this technique, the
athlete finally moves from bilateral (two-sided) exercises to unilateral (one-sided) work. This is an
important stage in the series. Working on half one-arm pushups will teach the athlete the balance
and positioning necessary to master full one-arm pushups. Because only one limb transmits the
moving forces, this exercise will also prepare the hand, wrist and shoulder joints for subsequent
steps. Half one-arm pushups are an essential exercise in the series, and must be mastered for the
reasons given above. However, the elbows are not required to bend very greatly in this exercise, so
it must never be performed by itself as the sum total of any pushup program. It should always be
followed with a movement where the elbows are bent beyond ninety degrees in the bottom position;
either close pushups or unevenpushups should be added afterwards.`,
				[3]standard{{}, {}, {}},
				``},
			{"Lever Pushups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"One-Arm Pushups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
		},
	},
	{
		"Squats",
		[]exInfo{
			{"Shoulderstand Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Jackknife Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Supported Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Half Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Full Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Close Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Uneven Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"1/2 One-Leg Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Assisted One-Leg Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"One-Leg Squats", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
		},
	},
	{
		"Pullups",
		[]exInfo{
			{"Vertical Pulls", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Horizontal Pulls", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Jackknife Pulls", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Half Pullups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Full Pullups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Close Pullups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Uneven Pullups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"1/2 One-Arm Pullups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Assisted One-Arm Pullups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"One-Arm Pullups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
		},
	},
	{
		"Leg Raises",
		[]exInfo{
			{"Knee Tucks", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Flat Knee Raises", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Flat Bent Leg Raises", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Flat Frog Raises", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Flat Straight Leg Raises", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Hanging Knee Raises", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Hanging Bent Leg Raises", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Hanging Frog Raises", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Partial Straight Leg Raises", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Hanging Straight Leg Raises", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
		},
	},
	{
		"Bridges",
		[]exInfo{
			{"Short Bridges", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Straight Bridges", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Angled Bridges", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Head Bridges", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Half Bridges", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Full Bridges", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Wall Walking Bridges (Down)", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Wall Walking Bridges (Up)", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Closing Bridges", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Stand-to-Stand Bridges", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
		},
	},
	{
		"Handstand Pushups",
		[]exInfo{
			{"Wall Headstands", "s", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Crow Stands", "s", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Wall Handstands", "s", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Half Handstand Pushups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Handstand Pushups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Close Handstand Pushups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Uneven Handstand Pushups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"1/2 One-Arm Handstand Pushups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"Lever Handstand Pushups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
			{"One-Arm Handstand Pushups", "", 0,
				``,
				``,
				[3]standard{{}, {}, {}},
				``},
		},
	},
}
