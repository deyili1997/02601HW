This is a bacteria-phage simulator.
This can simulate the bacteria-phage therapy to treat sepsis.

Input: 
plasmaSize: How large the plasma area is. Default: 600
bacNum: The number of initial uninfected bacteria in the plasma. Default: 500
phageNum: The number of initial phages in the plasma. Default: 300
immuneCellNum: The number of immune cells in the plasma. Default: 30
bacLifespan: How long a bacteria will live. Default: 600
bacReplicateRate: How fast a bacteria will replicate itself. The smaller the value is, the faster the bacteria replicate itself. Default: 200
lysisRate: How fast an infected bacteria will be lysed by the inside phage. The smaller the value is, the faster the bacteria being lysed. Default: 80
phageOffspring: The number of offsprings a phage can produce. Default: 4
numGens: The number of generations. Larger value will produce longer simulation. Default: 2000
(all integers) 

Output: 
An animation of “numGens ” generations.

Some basic rules for this physiological system:

1. Bacteria (green cycle):
(1) A bacteria has its lifespan as specified by the parameter “bacLifespan”, the brightness of the green circle indicate its lifespan. If it is dying, then it becomes darker and darker. If it is dead, then it disappears.
(2) A bacteria can replicate itself after some random rounds specified by “bacReplicateRate” parameter. The smaller “bacReplicateRate” is, the faster bacteria replicate themselves.

2. Bacteria-phage (small yellow dot):
(1) If a bacteria-phage meets a bacteria, then it will identify whether the bacteria can be infected. If the bacteria can be infected (according to the type of the bacteria), then the phage will enter into the bacteria and the bacteria will turn from green to orange indicating that this is a infected bacteria. If the bacteria cannot be infected, then it swims away.
(2) After a random rounds within “lysisRate”, the infected bacteria will be lysed and produce the number of “phageOffspring” children bacteria-phages.

3. Immune cell (big white cycle):
(1) Immune cells have specificity for bacteria or infected bacteria or bacteria-phage. That means they will track the organism they have specificity for. But they will eat whatever they encounter along the way.