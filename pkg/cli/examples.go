// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var genExamplesCmd = &cobra.Command{
	Use:   "example-data",
	Short: "generate example SQL code suitable for use with CockroachDB",
	Long: `This command generates example SQL code that shows various CockroachDB features and
is suitable to populate an example database for demonstration and education purposes.
The command takes an optional parameter which specifies which example to generate.
Available examples:
- startrek: a database containing a table of episodes and a table of
  quotes from the eponymous TV show.
- intro: a database containing a single table with a hidden message.
`,
	RunE: MaybeDecorateGRPCError(runGenExamplesCmd),
}

func runGenExamplesCmd(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return usageAndError(cmd)
	}

	example := "startrek"
	if len(args) > 0 {
		example = args[0]
	}

	if sql, ok := examples[example]; ok {
		fmt.Print(sql)
	} else {
		return fmt.Errorf("don't know how to generate example data for %q", example)
	}
	fmt.Println(footerComment)
	return nil
}

var examples = map[string]string{
	"intro":    introSQL,
	"startrek": startrekSQL,
}

const introSQL = `
CREATE DATABASE IF NOT EXISTS intro;
SET DATABASE=intro;
DROP TABLE IF EXISTS mytable;
CREATE TABLE mytable(l INT PRIMARY KEY, v TEXT);
INSERT INTO mytable(l,v) VALUES (0, X'215F5F61616177776D716D716D7777776161732C2C5F20202020202020202E5F5F6161617777776D716D716D77776161612C2C');
INSERT INTO mytable(l,v) VALUES (2, X'212256543F212222225E7E7E5E2222223F3F5424576D7161612C5F6175716D5742543F212222225E7E7E5E5E22223F3F59565E');
INSERT INTO mytable(l,v) VALUES (4, X'212020202020202020202020202020202020202020223F23236D5723233F222D20202020202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (6, X'21202043204F204E2047205220412054205320205F616D235A3F3F41236D612C2020202020202020202020592020202020202020');
INSERT INTO mytable(l,v) VALUES (8, X'2120202020202020202020202020202020205F756D6D5922202020202239236D612C2020202020202041202020202020202020');
INSERT INTO mytable(l,v) VALUES (10, X'2120202020202020202020202020202020766D235A28202020202020202029586D6D7320202020592020202020202020202020');
INSERT INTO mytable(l,v) VALUES (12, X'2120202020202020202020202020202E6A232323236D6D6D23232323236D6D236D2323362E2020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (14, X'2120202057204F20572021202020206A6D6D2323236D6D2323232323236D236D6D6D2323362020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (16, X'21202020202020202020202020205D236D652A586D236D236D6D23236D236D2323535823236320202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (18, X'2120202020202020202020202020646D237C7C2B2A2423236D236D6D236D235376766E23236D20202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (20, X'212020202020202020202020203A6D6D453D7C2B7C7C5323236D23236D23316E766E6E5823233B20202020204120202020202020');
INSERT INTO mytable(l,v) VALUES (22, X'212020202020202020202020203A6D23682B7C2B2B2B3D586D6D236D23316E766E6E76646D6D3B20202020204D202020202020');
INSERT INTO mytable(l,v) VALUES (24, X'212059202020202020202020202024236D3E2B7C2B7C7C7C23236D23316E766E6E6E6E6D6D2320202020202041202020202020');
INSERT INTO mytable(l,v) VALUES (26, X'2120204F202020202020202020205D23237A2B7C2B7C2B7C33236D456E6E6E6E766E642323662020202020205A202020202020');
INSERT INTO mytable(l,v) VALUES (28, X'212020205520204420202020202020342323637C2B7C2B7C5D6D236B766E766E6E6F2323502020202020202045202020202020');
INSERT INTO mytable(l,v) VALUES (30, X'2120202020202020492020202020202034236D612B7C2B2B5D6D6D68766E6E7671232350602020202020202021202020202020');
INSERT INTO mytable(l,v) VALUES (32, X'21202020202020202044204920202020203F242371252B7C646D6D6D766E6E6D23232120202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (34, X'2120202020202020202020205420202020202D3423237775236D6D237077232337272020202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (36, X'21202020202020202020202020202020202020202D3F2423236D23232323592720202020202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (38, X'21202020202020202020202020202121202020202020202259232359222D202020202020202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (40, X'21');
INSERT INTO mytable(l,v) VALUES (41, X'215F5F61616177776D716D716D7777776161732C2C5F20202020202020202E5F5F6161617777776D716D716D77776161612C2C');
INSERT INTO mytable(l,v) VALUES (39, X'212256543F212222225E7E7E5E2222223F3F5424576D7161612C5F6175716D5742543F212222225E7E7E5E5E22223F3F59565E');
INSERT INTO mytable(l,v) VALUES (37, X'212020202020202020202020202020202020202020223F23236D5723233F222D20202020202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (35, X'21202043204F204E2047205220412054205320205F616D235A3F3F41236D612C2020202020202020202020592020202020202020');
INSERT INTO mytable(l,v) VALUES (33, X'2120202020202020202020202020202020205F756D6D5922202020202239236D612C2020202020202041202020202020202020');
INSERT INTO mytable(l,v) VALUES (31, X'2120202020202020202020202020202020766D235A28202020202020202029586D6D7320202020592020202020202020202020');
INSERT INTO mytable(l,v) VALUES (29, X'2120202020202020202020202020202E6A232323236D6D6D23232323236D6D236D2323362E2020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (27, X'2120202057204F20572021202020206A6D6D2323236D6D2323232323236D236D6D6D2323362020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (25, X'21202020202020202020202020205D236D652A586D236D236D6D23236D236D2323535823236320202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (23, X'2120202020202020202020202020646D237C7C2B2A2423236D236D6D236D235376766E23236D20202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (21, X'212020202020202020202020203A6D6D453D7C2B7C7C5323236D23236D23316E766E6E5823233B20202020204120202020202020');
INSERT INTO mytable(l,v) VALUES (19, X'212020202020202020202020203A6D23682B7C2B2B2B3D586D6D236D23316E766E6E76646D6D3B20202020204D202020202020');
INSERT INTO mytable(l,v) VALUES (17, X'212059202020202020202020202024236D3E2B7C2B7C7C7C23236D23316E766E6E6E6E6D6D2320202020202041202020202020');
INSERT INTO mytable(l,v) VALUES (15, X'2120204F202020202020202020205D23237A2B7C2B7C2B7C33236D456E6E6E6E766E642323662020202020205A202020202020');
INSERT INTO mytable(l,v) VALUES (13, X'212020205520204420202020202020342323637C2B7C2B7C5D6D236B766E766E6E6F2323502020202020202045202020202020');
INSERT INTO mytable(l,v) VALUES (11, X'2120202020202020492020202020202034236D612B7C2B2B5D6D6D68766E6E7671232350602020202020202021202020202020');
INSERT INTO mytable(l,v) VALUES (9, X'21202020202020202044204920202020203F242371252B7C646D6D6D766E6E6D23232120202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (7, X'2120202020202020202020205420202020202D3423237775236D6D237077232337272020202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (5, X'21202020202020202020202020202020202020202D3F2423236D23232323592720202020202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (3, X'21202020202020202020202020202121202020202020202259232359222D202020202020202020202020202020202020202020');
INSERT INTO mytable(l,v) VALUES (1, X'21');
`

const startrekSQL = `
CREATE DATABASE IF NOT EXISTS startrek;
SET DATABASE=startrek;
DROP TABLE IF EXISTS quotes;
DROP TABLE IF EXISTS episodes;
CREATE TABLE episodes (id INT PRIMARY KEY, season INT, num INT, title TEXT, stardate DECIMAL);
-- The data that follows was derived from the 'startrek' fortune cookie file.
INSERT INTO episodes (id, season, num, title, stardate) VALUES
(1, 1, 1, 'The Man Trap', 1531.1),
(2, 1, 2, 'Charlie X', 1533.6),
(3, 1, 3, 'Where No Man Has Gone Before', 1312.4),
(4, 1, 4, 'The Naked Time', 1704.2),
(5, 1, 5, 'The Enemy Within', 1672.1),
(6, 1, 6, 'Mudd''s Women', 1329.8),
(7, 1, 7, 'What Are Little Girls Made Of?', 2712.4),
(8, 1, 8, 'Miri', 2713.5),
(9, 1, 9, 'Dagger of the Mind', 2715.1),
(10, 1, 10, 'The Corbomite Maneuver', 1512.2),
(11, 1, 11, 'The Menagerie, Part I', 3012.4),
(12, 1, 12, 'The Menagerie, Part II', 3013.1),
(13, 1, 13, 'The Conscience of the King', 2817.6),
(14, 1, 14, 'Balance of Terror', 1709.2),
(15, 1, 15, 'Shore Leave', 3025.3),
(16, 1, 16, 'The Galileo Seven', 2821.5),
(17, 1, 17, 'The Squire of Gothos', 2124.5),
(18, 1, 18, 'Arena', 3045.6),
(19, 1, 19, 'Tomorrow Is Yesterday', 3113.2),
(20, 1, 20, 'Court Martial', 2947.3),
(21, 1, 21, 'The Return of the Archons', 3156.2),
(22, 1, 22, 'Space Seed', 3141.9),
(23, 1, 23, 'A Taste of Armageddon', 3192.1),
(24, 1, 24, 'This Side of Paradise', 3417.3),
(25, 1, 25, 'The Devil in the Dark', 3196.1),
(26, 1, 26, 'Errand of Mercy', 3198.4),
(27, 1, 27, 'The Alternative Factor', 3087.6),
(28, 1, 28, 'The City on the Edge of Forever', 3134.0),
(29, 1, 29, 'Operation: Annihilate!', 3287.2),
(30, 2, 1, 'Amok Time', 3372.7),
(31, 2, 2, 'Who Mourns for Adonais?', 3468.1),
(32, 2, 3, 'The Changeling', 3541.9),
(33, 2, 4, 'Mirror, Mirror', NULL),
(34, 2, 5, 'The Apple', 3715.3),
(35, 2, 6, 'The Doomsday Machine', 4202.9),
(36, 2, 7, 'Catspaw', 3018.2),
(37, 2, 8, 'I, Mudd', 4513.3),
(38, 2, 9, 'Metamorphosis', 3219.4),
(39, 2, 10, 'Journey to Babel', 3842.3),
(40, 2, 11, 'Friday''s Child', 3497.2),
(41, 2, 12, 'The Deadly Years', 3478.2),
(42, 2, 13, 'Obsession', 3619.2),
(43, 2, 14, 'Wolf in the Fold', 3614.9),
(44, 2, 15, 'The Trouble with Tribbles', 4523.3),
(45, 2, 16, 'The Gamesters of Triskelion', 3211.8),
(46, 2, 17, 'A Piece of the Action', 4598.0),
(47, 2, 18, 'The Immunity Syndrome', 4307.1),
(48, 2, 19, 'A Private Little War', 4211.4),
(49, 2, 20, 'Return to Tomorrow', 4768.3),
(50, 2, 21, 'Patterns of Force', 2534.0),
(51, 2, 22, 'By Any Other Name', 4657.5),
(52, 2, 23, 'The Omega Glory', NULL),
(53, 2, 24, 'The Ultimate Computer', 4729.4),
(54, 2, 25, 'Bread and Circuses', 4040.7),
(55, 2, 26, 'Assignment: Earth', NULL),
(56, 3, 1, 'Spock''s Brain', 5431.4),
(57, 3, 2, 'The Enterprise Incident', 5027.3),
(58, 3, 3, 'The Paradise Syndrome', 4842.6),
(59, 3, 4, 'And the Children Shall Lead', 5029.5),
(60, 3, 5, 'Is There in Truth No Beauty?', 5630.7),
(61, 3, 6, 'Spectre of the Gun', 4385.3),
(62, 3, 7, 'Day of the Dove', 5630.3),
(63, 3, 8, 'For the World Is Hollow and I Have Touched the Sky', 5476.3),
(64, 3, 9, 'The Tholian Web', 5693.2),
(65, 3, 10, 'Plato''s Stepchildren', 5784.2),
(66, 3, 11, 'Wink of an Eye', 5710.5),
(67, 3, 12, 'The Empath', 5121.5),
(68, 3, 13, 'Elaan of Troyius', 4372.5),
(69, 3, 14, 'Whom Gods Destroy', 5718.3),
(70, 3, 15, 'Let That Be Your Last Battlefield', 5730.2),
(71, 3, 16, 'The Mark of Gideon', 5423.4),
(72, 3, 17, 'That Which Survives', NULL),
(73, 3, 18, 'The Lights of Zetar', 5725.3),
(74, 3, 19, 'Requiem for Methuselah', 5843.7),
(75, 3, 20, 'The Way to Eden', 5832.3),
(76, 3, 21, 'The Cloud Minders', 5818.4),
(77, 3, 22, 'The Savage Curtain', 5906.4),
(78, 3, 23, 'All Our Yesterdays', 5943.7),
(79, 3, 24, 'Turnabout Intruder', 5928.5);
CREATE TABLE quotes (quote TEXT, characters TEXT, stardate DECIMAL, episode INT REFERENCES episodes(id), INDEX(episode));
INSERT INTO quotes (quote, characters, stardate, episode) VALUES
('"... freedom ... is a worship word..." "It is our worship word too."', 'Cloud William and Kirk', NULL, 52),
('"Beauty is transitory." "Beauty survives."', 'Spock and Kirk', NULL, 72),
('"Can you imagine how life could be improved if we could do away with jealousy, greed, hate ..." "It can also be improved by eliminating love, tenderness, sentiment -- the other side of the coin"', 'Dr. Roger Corby and Kirk', 2712.4, 7),
('"Evil does seek to maintain power by suppressing the truth." "Or by misleading the innocent."', 'Spock and McCoy', 5029.5, 59),
('"Get back to your stations!" "We''re beaming down to the planet, sir."', 'Kirk and Mr. Leslie', 3417.3, 24),
('"I think they''re going to take all this money that we spend now on war and death --" "And make them spend it on life."', 'Edith Keeler and Kirk', NULL, 28),
('"It''s hard to believe that something which is neither seen nor felt can do so much harm." "That''s true.  But an idea can''t be seen or felt.  And that''s what kept the Troglytes in the mines all these centuries.  A mistaken idea."', 'Vanna and Kirk', 5819.0, 76),
('"Life and death are seldom logical." "But attaining a desired goal always is."', 'McCoy and Spock', 2821.7, 16),
('"Logic and practical information do not seem to apply here." "You admit that?" "To deny the facts would be illogical, Doctor"', 'Spock and McCoy', NULL, 46),
('"No one talks peace unless he''s ready to back it up with war." "He talks of peace if it is the only way to live."', 'Colonel Green and Surak of Vulcan', 5906.5, 77),
('"That unit is a woman." "A mass of conflicting impulses."', 'Spock and Nomad', 3541.9, 32),
('"The combination of a number of things to make existence worthwhile." "Yes, the philosophy of ''none,'' meaning ''all.''"', 'Spock and Lincoln', 5906.4, 77),
('"The glory of creation is in its infinite diversity." "And in the way our differences combine to create meaning and beauty."', 'Dr. Miranda Jones and Spock', 5630.8, 60),
('"The release of emotion is what keeps us healthy.  Emotionally healthy." "That may be, Doctor.  However, I have noted that the healthy release of emotion is frequently unhealthy for those closest to you."', 'McCoy and Spock', 5784.3, 65),
('"There''s only one kind of woman ..." "Or man, for that matter.  You either believe in yourself or you don''t."', 'Kirk and Harry Mudd', 1330.1, 6),
('"We have the right to survive!" "Not by killing others."', 'Deela and Kirk', 5710.5, 66),
('"What a terrible way to die." "There are no good ways."', 'Sulu and Kirk', NULL, 72),
('"What happened to the crewman?" "The M-5 computer needed a new power source, the crewman merely got in the way."', 'Kirk and Dr. Richard Daystrom', 4731.3, 53),
('... bacteriological warfare ... hard to believe we were once foolish enough to play around with that.', 'McCoy', NULL, 52),
('... The prejudices people feel about each other disappear when they get to know each other.', 'Kirk', 4372.5, 68),
('... The things love can drive a man to -- the ecstasies, the miseries, the broken rules, the desperate chances, the glorious failures and the glorious victories.', 'McCoy', 5843.7, 74),
('A father doesn''t destroy his children.', 'Lt. Carolyn Palamas', 3468.1, 31),
('A little suffering is good for the soul.', 'Kirk', 1514.0, 10),
('A man either lives life as it happens to him, meets it head-on and licks it, or he turns his back on it and starts to wither away.', 'Dr. Boyce', NULL, 11),
('A princess should not be afraid -- not with a brave knight to protect her.', 'McCoy', 3025.3, 15),
('A star captain''s most solemn oath is that he will give his life, even his entire crew, rather than violate the Prime Directive.', 'Kirk', NULL, 52),
('A Vulcan can no sooner be disloyal than he can exist without breathing.', 'Kirk', 3012.4, 11),
('A woman should have compassion.', 'Kirk', 3018.2, 36),
('Actual war is a very messy business.  Very, very messy business.', 'Kirk', 3193.0, 23),
('After a time, you may find that "having" is not so pleasing a thing, after all, as "wanting."  It is not logical, but it is often true.', 'Spock', 3372.7, 30),
('All your people must learn before you can reach for the stars.', 'Kirk', 3259.2, 45),
('Another Armenia, Belgium ... the weak innocents who always seem to be located on a natural invasion route.', 'Kirk', 3198.4, 26),
('Another dream that failed.  There''s nothing sadder.', 'Kirk', 3417.3, 24),
('Another war ... must it always be so?  How many comrades have we lost in this way? ...  Obedience.  Duty.  Death, and more death ...', 'Romulan Commander', 1709.2, 14),
('Behind every great man, there is a woman -- urging him on.', 'Harry Mudd', 4513.3, 37),
('Blast medicine anyway!  We''ve learned to tie into every organ in the human body but one.  The brain!  The brain is what life is all about.', 'McCoy', 3012.4, 11),
('But it''s real.  And if it''s real it can be affected ...  we may not be able to break it, but, I''ll bet you credits to Navy Beans we can put a dent in it.', 'deSalle', 3018.2, 36),
('Change is the essential process of all existence.', 'Spock', 5730.2, 70),
('Compassion -- that''s the one thing no machine ever had.  Maybe it''s the one thing that keeps men ahead of them.', 'McCoy', 4731.3, 53),
('Computers make excellent and efficient servants, but I have no wish to serve under them.  Captain, a starship also runs on loyalty to one man.  And nothing can replace it or him.', 'Spock', 4729.4, 53),
('Conquest is easy. Control is not.', 'Kirk', NULL, 33),
('Death.  Destruction.  Disease.  Horror.  That''s what war is all about. That''s what makes it a thing to be avoided.', 'Kirk', 3193.0, 23),
('Death, when unnecessary, is a tragic thing.', 'Flint', 5843.7, 74),
('Do you know about being with somebody?  Wanting to be?  If I had the whole universe, I''d give it to you, Janice.  When I see you, I feel like I''m hungry all over.  Do you know how that feels?', 'Charlie Evans', 1535.8, 2),
('Do you know the one -- "All I ask is a tall ship, and a star to steer her by ..."  You could feel the wind at your back, about you ...  the sounds of the sea beneath you.  And even if you take away the wind and the water, it''s still the same.  The ship is yours ... you can feel her ... and the stars are still there.', 'Kirk', 4729.4, 53),
('[Doctors and Bartenders], We both get the same two kinds of customers -- the living and the dying.', 'Dr. Boyce', NULL, 11),
('Each kiss is as the first.', 'Miramanee, Kirk''s wife', 4842.6, 58),
('Earth -- mother of the most beautiful women in the universe.', 'Apollo', 3468.1, 31),
('Either one of us, by himself, is expendable.  Both of us are not.', 'Kirk', 3196.1, 25),
('Emotions are alien to me.  I''m a scientist.', 'Spock', 3417.3, 24),
('Even historians fail to learn from history -- they repeat the same mistakes.', 'John Gill', 2534.7, 50),
('Every living thing wants to survive.', 'Spock', 4731.3, 53),
('Extreme feminine beauty is always disturbing.', 'Spock', 5818.4, 76),
('Fascinating, a totally parochial attitude.', 'Spock', 3219.8, 38),
('Fascinating is a word I use for the unexpected.', 'Spock', 2124.5, 17),
('First study the enemy.  Seek weakness.', 'Romulan Commander', 1709.2, 14),
('Four thousand throats may be cut in one night by a running man.', 'Klingon Soldier', NULL, 62),
('Genius doesn''t work on an assembly line basis.  You can''t simply say, "Today I will be brilliant."', 'Kirk', 4731.3, 53),
('He''s dead, Jim', 'McCoy', 3196.1, 25),
('History tends to exaggerate.', 'Col. Green', 5906.4, 77),
('Humans do claim a great deal for that particular emotion (love).', 'Spock', 5725.6, 73),
('I am pleased to see that we have differences.  May we together become greater than the sum of both of us.', 'Surak of Vulcan', 5906.4, 77),
('I have never understood the female capacity to avoid a direct answer to any question.', 'Spock', 3417.3, 24),
('I object to intellect without discipline;  I object to power without constructive purpose.', 'Spock', 2124.5, 17),
('I realize that command does have its fascination, even under circumstances such as these, but I neither enjoy the idea of command nor am I frightened of it.  It simply exists, and I will do whatever logically needs to be done.', 'Spock', 2812.7, 16),
('I thought my people would grow tired of killing.  But you were right, they see it is easier than trading.  And it has its pleasures.  I feel it myself.  Like the hunt, but with richer rewards.', 'Apella', 4211.8, 48),
('If a man had a child who''d gone anti-social, killed perhaps, he''d still tend to protect that child.', 'McCoy', 4731.3, 53),
('If I can have honesty, it''s easier to overlook mistakes.', 'Kirk', 3141.9, 22),
('If some day we are defeated, well, war has its fortunes, good and bad.', 'Commander Kor', 3201.7, 26),
('If there are self-made purgatories, then we all have to live in them.', 'Spock', 3417.7, 24),
('I''m a soldier, not a diplomat.  I can only tell the truth.', 'Kirk', 3198.9, 26),
('I''m frequently appalled by the low regard you Earthmen have for life.', 'Spock', 2822.3, 16),
('Immortality consists largely of boredom.', 'Zefrem Cochrane', 3219.8, 38),
('In the strict scientific sense we all feed on death -- even vegetarians.', 'Spock', 3615.4, 43),
('Insufficient facts always invite danger.', 'Spock', 3141.9, 22),
('Insults are effective only where emotion is present.', 'Spock', 3468.1, 31),
('Intuition, however illogical, is recognized as a command prerogative.', 'Kirk', 3620.7, 42),
('Is not that the nature of men and women -- that the pleasure is in the learning of each other?', 'Natira, the High Priestess of Yonada', 5476.3, 63),
('Is truth not truth for all?', 'Natira', 5476.4, 63),
('It [being a Vulcan] means to adopt a philosophy, a way of life which is logical and beneficial.  We cannot disregard that philosophy merely for personal gain, no matter how important that gain might be.', 'Spock', 3842.4, 39),
('It is a human characteristic to love little animals, especially if they''re attractive in some way.', 'McCoy', 4525.6, 44),
('It is more rational to sacrifice one life than six.', 'Spock', 2822.3, 16),
('It is necessary to have purpose.', 'Alice #1', 4513.3, 37),
('It is undignified for a woman to play servant to a man who is not hers.', 'Spock', 3372.7, 30),
('It would be illogical to assume that all conditions remain stable.', 'Spock', 5027.3, 57),
('It would be illogical to kill without reason', 'Spock', 3842.4, 39),
('It would seem that evil retreats when forcibly confronted', 'Yarnek of Excalbia', 5906.5, 77),
('I''ve already got a female to worry about.  Her name is the Enterprise.', 'Kirk', 1514.0, 10),
('Killing is stupid; useless!', 'McCoy', 4211.8, 48),
('Killing is wrong.', 'Losira', NULL, 72),
('Knowledge, sir, should be free to all!', 'Harry Mudd', 4513.3, 37),
('Landru! Guide us!', 'A Beta 3-oid', 3157.4, 21),
('Leave bigotry in your quarters; there''s no room for it on the bridge.', 'Kirk', 1709.2, 14),
('Live long and prosper.', 'Spock', 3372.7, 30),
('Lots of people drink from the wrong bottle sometimes.', 'Edith Keeler', NULL, 28),
('Love sometimes expresses itself in sacrifice.', 'Kirk', 3220.3, 38),
('Madness has no purpose.  Or reason.  But it may have a goal.', 'Spock', 3088.7, 27),
('Many Myths are based on truth', 'Spock', 5832.3, 75),
('Men don''t talk peace unless they''re ready to back it up with war.', 'Col. Green', 5906.4, 77),
('Men of peace usually are [brave].', 'Spock', 5906.5, 77),
('Men will always be men -- no matter where they are.', 'Harry Mudd', 1329.8, 6),
('Military secrets are the most fleeting of all.', 'Spock', 5027.4, 57),
('Most legends have their basis in facts.', 'Kirk', 5029.5, 59),
('Murder is contrary to the laws of man and God.', 'M-5 Computer', 4731.3, 53),
('No more blah, blah, blah!', 'Kirk', 2713.6, 8),
('No one can guarantee the actions of another.', 'Spock', NULL, 62),
('No one may kill a man.  Not for any purpose.  It cannot be condoned.', 'Kirk', 5431.6, 56),
('No one wants war.', 'Kirk', 3201.7, 26),
('No problem is insoluble.', 'Dr. Janet Wallace', 3479.4, 41),
('Not one hundred percent efficient, of course ... but nothing ever is.', 'Kirk', 3219.8, 38),
('Oblivion together does not frighten me, beloved.', 'Thalassa (in Anne Mulhall''s body)', 4770.3, 49),
('Oh, that sound of male ego.  You travel halfway across the galaxy and it''s still the same song.', 'Eve McHuron', 1330.1, 6),
('On my planet, to rest is to rest -- to cease using energy.  To me, it is quite illogical to run up and down on green grass, using energy, instead of saving it.', 'Spock', 3025.2, 15),
('One does not thank logic.', 'Sarek', 3842.4, 39),
('One of the advantages of being a captain is being able to ask for advice without necessarily having to take it.', 'Kirk', 2715.2, 9),
('Only a fool fights in a burning house.', 'Kang the Klingon', NULL, 62),
('Our missions are peaceful -- not for conquest.  When we do battle, it is only because we have no choice.', 'Kirk', 2124.5, 17),
('Our way is peace.', 'Septimus, the Son Worshiper', 4040.7, 54),
('Pain is a thing of the mind.  The mind can be controlled.', 'Spock', 3287.2, 29),
('Peace was the way.', 'Kirk', NULL, 28),
('Power is danger.', 'The Centurion', 1709.2, 14),
('Prepare for tomorrow -- get ready.', 'Edith Keeler', NULL, 28),
('Punishment becomes ineffective after a certain point.  Men become insensitive.', 'Eneg', 2534.7, 50),
('Respect is a rational process', 'McCoy', 2822.3, 16),
('Romulan women are not like Vulcan females.  We are not dedicated to pure logic and the sterility of non-emotion.', 'Romulan Commander', 5027.3, 57),
('Schshschshchsch.', 'The Gorn', 3046.2, 18),
('Sometimes a feeling is all we humans have to go on.', 'Kirk', 3193.9, 23),
('Sometimes a man will tell his bartender things he''ll never tell his doctor.', 'Dr. Phillip Boyce', NULL, 11),
('Suffocating together ... would create heroic camaraderie.', 'Khan Noonian Singh', 3142.8, 22),
('Superior ability breeds superior ambition.', 'Spock', 3141.9, 22),
('The face of war has never changed.  Surely it is more logical to heal than to kill.', 'Surak of Vulcan', 5906.5, 77),
('The games have always strengthened us.  Death becomes a familiar pattern.  We don''t fear it as you do.', 'Proconsul Marcus Claudius', 4041.2, 54),
('The heart is not a logical organ.', 'Dr. Janet Wallace', 3479.4, 41),
('The idea of male and female are universal constants.', 'Kirk', 3219.8, 38),
('The joys of love made her human and the agonies of love destroyed her.', 'Spock', 5842.8, 74),
('The man on tops walks a lonely street; the "chain" of command is often a noose.', 'McCoy', 2818.9, 13),
('The more complex the mind, the greater the need for the simplicity of play.', 'Kirk', 3025.8, 15),
('The only solution is ... a balance of power.  We arm our side with exactly that much more.  A balance of power -- the trickiest, most difficult, dirtiest game of them all.  But the only one that preserves both sides.', 'Kirk', 4211.8, 48),
('The people of Gideon have always believed that life is sacred.  That the love of life is the greatest gift ... We are incapable of destroying or interfering with the creation of that which we love so deeply -- life in every form from fetus to developed being.', 'Hodin of Gideon', 5423.4, 71),
('The sight of death frightens them [Earthers].', 'Kras the Klingon', 3497.2, 40),
('The sooner our happiness together begins, the longer it will last.', 'Miramanee', 4842.6, 58),
('There are always alternatives.', 'Spock', 2822.3, 16),
('There are certain things men must do to remain men.', 'Kirk', 4929.4, 53),
('There are some things worth dying for.', 'Kirk', 3201.7, 26),
('There comes to all races an ultimate crisis which you have yet to face .... One day our minds became so powerful we dared think of ourselves as gods.', 'Sargon', 4768.3, 49),
('There is a multi-legged creature crawling on your shoulder.', 'Spock', 3193.9, 23),
('There is an old custom among my people.  When a woman saves a man''s life, he is grateful.', 'Nona, the Kanuto witch woman', 4211.8, 48),
('There is an order of things in this universe.', 'Apollo', 3468.1, 31),
('There''s a way out of any cage.', 'Captain Christopher Pike', NULL, 11),
('There''s another way to survive.  Mutual trust -- and help.', 'Kirk', NULL, 62),
('There''s no honorable way to kill, no gentle way to destroy.  There is nothing good in war.  Except its ending.', 'Abraham Lincoln', 5906.5, 77),
('There''s nothing disgusting about it [the Companion].  It''s just another life form, that''s all.  You get used to those things.', 'McCoy', 3219.8, 38),
('This cultural mystique surrounding the biological function -- you realize humans are overly preoccupied with the subject.', 'Kelinda the Kelvan', 4658.9, 51),
('Those who hate and fight must stop themselves -- otherwise it is not stopped.', 'Spock', NULL, 62),
('Time is fluid ... like a river with currents, eddies, backwash.', 'Spock', 3134.0, 28),
('To live is always desirable.', 'Eleen the Capellan', 3498.9, 40),
('Too much of anything, even love, isn''t necessarily a good thing.', 'Kirk', 4525.6, 44),
('Totally illogical, there was no chance.', 'Spock', 2822.3, 16),
('Uncontrolled power will turn even saints into savages.  And we can all be counted on to live down to our lowest impulses.', 'Parmen', 5784.3, 65),
('Violence in reality is quite different from theory.', 'Spock', 5818.4, 76),
('Virtue is a relative term.', 'Spock', 3499.1, 40),
('Vulcans believe peace should not depend on force.', 'Amanda', 3842.3, 39),
('Vulcans do not approve of violence.', 'Spock', 3842.4, 39),
('Vulcans never bluff.', 'Spock', 4202.1, 35),
('Vulcans worship peace above all.', 'McCoy', 4768.3, 49),
('Wait!  You have not been prepared!', 'Mr. Atoz', 3113.2, 19),
('[War] is instinctive.  But the instinct can be fought.  We''re human beings with the blood of a million savage years on our hands!  But we can stop it.  We can admit that we''re killers ... but we''re not going to kill today.  That''s all it takes!  Knowing that we''re not going to kill today!', 'Kirk', 3193.0, 23),
('War is never imperative.', 'McCoy', 1709.2, 14),
('War isn''t a good life, but it''s life.', 'Kirk', 4211.8, 48),
('We do not colonize.  We conquer.  We rule.  There is no other way for us.', 'Rojan', 4657.5, 51),
('We fight only when there is no other choice.  We prefer the ways of peaceful contact.', 'Kirk', 4385.3, 61),
('We have found all life forms in the galaxy are capable of superior development.', 'Kirk', 3211.7, 45),
('We have phasers, I vote we blast ''em!', 'Bailey', 1514.2, 10),
('We Klingons believe as you do -- the sick should die.  Only the strong should live.', 'Kras', 3497.2, 40),
('We''re all sorry for the other guy when he loses his job to a machine. But when it comes to your job -- that''s different.  And it always will be different.', 'McCoy', 4729.4, 53),
('What kind of love is that?  Not to be loved; never to have shown love.', 'Commissioner Nancy Hedford', 3219.8, 38),
('When a child is taught ... it''s programmed with simple instructions -- and at some point, if its mind develops properly, it exceeds the sum of what it was taught, thinks independently.', 'Dr. Richard Daystrom', 4731.3, 53),
('When dreams become more important than reality, you give up travel, building, creating; you even forget how to repair the machines left behind by your ancestors.  You just sit living and reliving other lives left behind in the thought records.', 'Vina', NULL, 11),
('Where there''s no emotion, there''s no motive for violence.', 'Spock', 2715.1, 9),
('Witch!  Witch!  They''ll burn ya!', 'Hag', NULL, 19),
('Without facts, the decision cannot be made logically.  You must rely on your human intuition.', 'Spock', NULL, 55),
('Without followers, evil cannot spread.', 'Spock', 5029.5, 59),
('Without freedom of choice there is no creativity.', 'Kirk', 3157.4, 21),
('Women are more easily and more deeply terrified ... generating more sheer horror than the male of the species.', 'Spock', 3615.4, 43),
('Women professionals do tend to over-compensate.', 'Dr. Elizabeth Dehner', 1312.9, 3),
('Worlds may change, galaxies disintegrate, but a woman always remains a woman.', 'Kirk', 2818.9, 13),
('Yes, it is written.  Good shall always destroy evil.', 'Sirah the Yang', NULL, 52),
('You!  What PLANET is this?!', 'McCoy', 3134.0, 28),
('You are an excellent tactician, Captain.  You let your second in command attack while you sit and watch for weakness.', 'Khan Noonian Singh', 3141.9, 22),
('You can''t evaluate a man by logic alone.', 'McCoy', 4513.3, 37),
('You Earth people glorified organized violence for forty centuries.  But you imprison those who employ it privately.', 'Spock', 2715.1, 9),
('You go slow, be gentle.  It''s no one-way street -- you know how you feel and that''s all.  It''s how the girl feels too.  Don''t press.  If the girl feels anything for you at all, you''ll know.', 'Kirk', 1535.8, 2),
('You humans have that emotional need to express gratitude.  "You''re welcome," I believe, is the correct response.', 'Spock', 4041.2, 54),
('You say you are lying.  But if everything you say is a lie, then you are telling the truth.  You cannot tell the truth because everything you say is a lie.  You lie, you tell the truth ... but you cannot, for you lie.', 'Norman the android', 4513.3, 37),
('You speak of courage.  Obviously you do not know the difference between courage and foolhardiness.  Always it is the brave ones who die, the soldiers.', 'Kor, the Klingon Commander', 3201.7, 26),
('You''ll learn something about men and women -- the way they''re supposed to be.  Caring for each other, being happy with each other, being good to each other.  That''s what we call love.  You''ll like that a lot.', 'Kirk', 3715.6, 34),
('You''re dead, Jim.', 'McCoy', 3372.7, 30),
('You''re dead, Jim.', 'McCoy', NULL, 64),
('You''re too beautiful to ignore.  Too much woman.', 'Kirk to Yeoman Rand', NULL, 5),
('Youth doesn''t excuse everything.', 'Dr. Janice Lester (in Kirk''s body)', 5928.5, 79);
`

const footerComment = `--
--
-- If you can see this message, you probably want to redirect the output of
-- 'cockroach gen example-data' to a file, or pipe it as input to 'cockroach sql'.
`
