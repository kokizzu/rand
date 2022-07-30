// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE-go file.

package rand

import (
	"math"
)

/*
 * Exponential distribution
 *
 * See "The Ziggurat Method for Generating Random Variables"
 * (Marsaglia & Tsang, 2000)
 * https://www.jstatsoft.org/v05/i08/paper [pdf]
 *
 * Fixed correlation, see https://github.com/flyingmutant/rand/issues/3
 */

const (
	re = 7.69711747013104972
)

// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1).
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
//
//  sample = ExpFloat64() / desiredRateParameter
//
func (r *Rand) ExpFloat64() float64 {
	for {
		v := r.Uint64()
		j := uint64(uint32(v >> 8))
		i := v & 0xFF
		x := float64(j) * we[i]
		if j < ke[i] {
			return x
		}
		if i == 0 {
			return re - math.Log(r.Float64())
		}
		if fe[i]+r.Float64()*(fe[i-1]-fe[i]) < math.Exp(-x) {
			return x
		}
	}
}

var ke = [256]uint64{
	0xe290a139, 0x0, 0x9beadebc, 0xc377ac71, 0xd4ddb990,
	0xde893fb8, 0xe4a8e87c, 0xe8dff16a, 0xebf2deab, 0xee49a6e8,
	0xf0204efd, 0xf19bdb8e, 0xf2d458bb, 0xf3da104b, 0xf4b86d78,
	0xf577ad8a, 0xf61de83d, 0xf6afb784, 0xf730a573, 0xf7a37651,
	0xf80a5bb6, 0xf867189d, 0xf8bb1b4f, 0xf9079062, 0xf94d70ca,
	0xf98d8c7d, 0xf9c8928a, 0xf9ff175b, 0xfa319996, 0xfa6085f8,
	0xfa8c3a62, 0xfab5084e, 0xfadb36c8, 0xfaff0410, 0xfb20a6ea,
	0xfb404fb4, 0xfb5e2951, 0xfb7a59e9, 0xfb95038c, 0xfbae44ba,
	0xfbc638d8, 0xfbdcf892, 0xfbf29a30, 0xfc0731df, 0xfc1ad1ed,
	0xfc2d8b02, 0xfc3f6c4d, 0xfc5083ac, 0xfc60ddd1, 0xfc708662,
	0xfc7f8810, 0xfc8decb4, 0xfc9bbd62, 0xfca9027c, 0xfcb5c3c3,
	0xfcc20864, 0xfccdd70a, 0xfcd935e3, 0xfce42ab0, 0xfceebace,
	0xfcf8eb3b, 0xfd02c0a0, 0xfd0c3f59, 0xfd156b7b, 0xfd1e48d6,
	0xfd26daff, 0xfd2f2552, 0xfd372af7, 0xfd3eeee5, 0xfd4673e7,
	0xfd4dbc9e, 0xfd54cb85, 0xfd5ba2f2, 0xfd62451b, 0xfd68b415,
	0xfd6ef1da, 0xfd750047, 0xfd7ae120, 0xfd809612, 0xfd8620b4,
	0xfd8b8285, 0xfd90bcf5, 0xfd95d15e, 0xfd9ac10b, 0xfd9f8d36,
	0xfda43708, 0xfda8bf9e, 0xfdad2806, 0xfdb17141, 0xfdb59c46,
	0xfdb9a9fd, 0xfdbd9b46, 0xfdc170f6, 0xfdc52bd8, 0xfdc8ccac,
	0xfdcc542d, 0xfdcfc30b, 0xfdd319ef, 0xfdd6597a, 0xfdd98245,
	0xfddc94e5, 0xfddf91e6, 0xfde279ce, 0xfde54d1f, 0xfde80c52,
	0xfdeab7de, 0xfded5034, 0xfdefd5be, 0xfdf248e3, 0xfdf4aa06,
	0xfdf6f984, 0xfdf937b6, 0xfdfb64f4, 0xfdfd818d, 0xfdff8dd0,
	0xfe018a08, 0xfe03767a, 0xfe05536c, 0xfe07211c, 0xfe08dfc9,
	0xfe0a8fab, 0xfe0c30fb, 0xfe0dc3ec, 0xfe0f48b1, 0xfe10bf76,
	0xfe122869, 0xfe1383b4, 0xfe14d17c, 0xfe1611e7, 0xfe174516,
	0xfe186b2a, 0xfe19843e, 0xfe1a9070, 0xfe1b8fd6, 0xfe1c8289,
	0xfe1d689b, 0xfe1e4220, 0xfe1f0f26, 0xfe1fcfbc, 0xfe2083ed,
	0xfe212bc3, 0xfe21c745, 0xfe225678, 0xfe22d95f, 0xfe234ffb,
	0xfe23ba4a, 0xfe241849, 0xfe2469f2, 0xfe24af3c, 0xfe24e81e,
	0xfe25148b, 0xfe253474, 0xfe2547c7, 0xfe254e70, 0xfe25485a,
	0xfe25356a, 0xfe251586, 0xfe24e88f, 0xfe24ae64, 0xfe2466e1,
	0xfe2411df, 0xfe23af34, 0xfe233eb4, 0xfe22c02c, 0xfe22336b,
	0xfe219838, 0xfe20ee58, 0xfe20358c, 0xfe1f6d92, 0xfe1e9621,
	0xfe1daef0, 0xfe1cb7ac, 0xfe1bb002, 0xfe1a9798, 0xfe196e0d,
	0xfe1832fd, 0xfe16e5fe, 0xfe15869d, 0xfe141464, 0xfe128ed3,
	0xfe10f565, 0xfe0f478c, 0xfe0d84b1, 0xfe0bac36, 0xfe09bd73,
	0xfe07b7b5, 0xfe059a40, 0xfe03644c, 0xfe011504, 0xfdfeab88,
	0xfdfc26e9, 0xfdf98629, 0xfdf6c83b, 0xfdf3ec01, 0xfdf0f04a,
	0xfdedd3d1, 0xfdea953d, 0xfde7331e, 0xfde3abe9, 0xfddffdfb,
	0xfddc2791, 0xfdd826cd, 0xfdd3f9a8, 0xfdcf9dfc, 0xfdcb1176,
	0xfdc65198, 0xfdc15bb3, 0xfdbc2ce2, 0xfdb6c206, 0xfdb117be,
	0xfdab2a63, 0xfda4f5fd, 0xfd9e7640, 0xfd97a67a, 0xfd908192,
	0xfd8901f2, 0xfd812182, 0xfd78d98e, 0xfd7022bb, 0xfd66f4ed,
	0xfd5d4732, 0xfd530f9c, 0xfd48432b, 0xfd3cd59a, 0xfd30b936,
	0xfd23dea4, 0xfd16349e, 0xfd07a7a3, 0xfcf8219b, 0xfce7895b,
	0xfcd5c220, 0xfcc2aadb, 0xfcae1d5e, 0xfc97ed4e, 0xfc7fe6d4,
	0xfc65ccf3, 0xfc495762, 0xfc2a2fc8, 0xfc07ee19, 0xfbe213c1,
	0xfbb8051a, 0xfb890078, 0xfb5411a5, 0xfb180005, 0xfad33482,
	0xfa839276, 0xfa263b32, 0xf9b72d1c, 0xf930a1a2, 0xf889f023,
	0xf7b577d2, 0xf69c650c, 0xf51530f0, 0xf2cb0e3c, 0xeeefb15d,
	0xe6da6ecf,
}
var we = [256]float64{
	2.0249554585039273e-09, 1.486674039973914e-11, 2.4409617196260667e-11,
	3.196880708914543e-11, 3.844677064665297e-11, 4.422820397243645e-11,
	4.9516444707048736e-11, 5.443358865093317e-11, 5.905944001532905e-11,
	6.34494203791173e-11, 6.764381087646596e-11, 7.167294497483693e-11,
	7.5560323199469e-11, 7.932458097693723e-11, 8.298078557904666e-11,
	8.654132143825228e-11, 9.001651265218844e-11, 9.341507193080098e-11,
	9.674443155535412e-11, 1.0001099208030168e-10, 1.0322031240760172e-10,
	1.0637725725104571e-10, 1.0948611308871047e-10, 1.125506804449162e-10,
	1.1557434814019853e-10, 1.1856015362861901e-10, 1.2151083247552976e-10,
	1.2442885926858652e-10, 1.2731648170466315e-10, 1.3017574919190738e-10,
	1.3300853700670142e-10, 1.3581656682043557e-10, 1.3860142424039144e-10,
	1.41364573878306e-10, 1.44107372359111e-10, 1.4683107960351982e-10,
	1.4953686865617902e-10, 1.522258342820371e-10, 1.5489900051445652e-10,
	1.575573273071839e-10, 1.6020171641692235e-10, 1.6283301662263273e-10,
	1.654520283708477e-10, 1.680595079224455e-10, 1.7065617106490894e-10,
	1.7324269644462226e-10, 1.7581972856586386e-10, 1.783878804965491e-10,
	1.8094773631522658e-10, 1.8349985332914917e-10, 1.8604476408927868e-10,
	1.8858297822471202e-10, 1.9111498411614723e-10, 1.9364125042554764e-10,
	1.9616222749705624e-10, 1.986783486423952e-10, 2.011900313224188e-10,
	2.0369767823513247e-10, 2.0620167831931063e-10, 2.087024076818232e-10,
	2.112002304558852e-10, 2.1369549959666193e-10, 2.1618855761997644e-10,
	2.186797372892644e-10, 2.211693622553898e-10, 2.2365774765346816e-10,
	2.261452006604297e-10, 2.2863202101668864e-10, 2.3111850151495907e-10,
	2.3360492845897016e-10, 2.360915820945744e-10, 2.3857873701551393e-10,
	2.4106666254590454e-10, 2.4355562310131354e-10, 2.460458785301426e-10,
	2.485376844368799e-10, 2.510312924886523e-10, 2.535269507063894e-10,
	2.5602490374180415e-10, 2.585253931412964e-10, 2.6102865759779937e-10,
	2.635349331915095e-10, 2.6604445362036876e-10, 2.68557450421102e-10,
	2.710741531815563e-10, 2.7359478974503266e-10, 2.7611958640725403e-10,
	2.786487681065693e-10, 2.8118255860795293e-10, 2.837211806813232e-10,
	2.8626485627467025e-10, 2.8881380668245404e-10, 2.913682527097064e-10,
	2.9392841483224535e-10, 2.9649451335338897e-10, 2.9906676855753465e-10,
	3.016454008609524e-10, 3.042306309601231e-10, 3.068226799779396e-10,
	3.0942176960807203e-10, 3.120281222577915e-10, 3.146419611895305e-10,
	3.172635106614526e-10, 3.1989299606729535e-10, 3.225306440757405e-10,
	3.2517668276956344e-10, 3.278313417848049e-10, 3.304948524502066e-10,
	3.33167447927147e-10, 3.358493633503123e-10, 3.3854083596933475e-10,
	3.412421052916313e-10, 3.4395341322667273e-10, 3.46675004231917e-10,
	3.494071254606396e-10, 3.5215002691189675e-10, 3.549039615828604e-10,
	3.5766918562376677e-10, 3.604459584957252e-10, 3.6323454313163823e-10,
	3.660352061004911e-10, 3.688482177752741e-10, 3.716738525048091e-10,
	3.7451238878976035e-10, 3.7736410946311836e-10, 3.8022930187545505e-10,
	3.831082580852609e-10, 3.860012750546849e-10, 3.8890865485101284e-10,
	3.918307048542317e-10, 3.947677379710454e-10, 3.977200728557206e-10,
	4.0068803413816137e-10, 4.0367195265962996e-10, 4.066721657165498e-10,
	4.096890173128513e-10, 4.127228584213427e-10, 4.157740472546139e-10,
	4.1884294954600986e-10, 4.2192993884123634e-10, 4.2503539680119584e-10,
	4.281597135166822e-10, 4.3130328783559965e-10, 4.3446652770341084e-10,
	4.3764985051756043e-10, 4.408536834966642e-10, 4.4407846406530283e-10,
	4.4732464025531137e-10, 4.505926711245093e-10, 4.538830271938779e-10,
	4.5719619090425494e-10, 4.6053265709368507e-10, 4.63892933496641e-10,
	4.672775412664091e-10, 4.706870155220212e-10, 4.74121905921206e-10,
	4.775827772609386e-10, 4.810702101072703e-10, 4.845848014562447e-10,
	4.881271654278306e-10, 4.916979339949417e-10, 4.952977577497641e-10,
	4.989273067097741e-10, 5.025872711660073e-10, 5.062783625763314e-10,
	5.100013145066842e-10, 5.137568836234657e-10, 5.175458507405212e-10,
	5.21369021924424e-10, 5.252272296620576e-10, 5.291213340948229e-10,
	5.330522243241474e-10, 5.370208197933575e-10, 5.410280717513981e-10,
	5.4507496480435e-10, 5.491625185611978e-10, 5.532917893808662e-10,
	5.574638722281573e-10, 5.616799026468934e-10, 5.659410588593269e-10,
	5.702485640016967e-10, 5.746036885067274e-10, 5.790077526448783e-10,
	5.834621292372686e-10, 5.879682465544501e-10, 5.92527591416582e-10,
	5.971417125121004e-10, 6.018122239536933e-10, 6.065408090923064e-10,
	6.11329224612049e-10, 6.161793049312685e-10, 6.21092966937755e-10,
	6.260722150890632e-10, 6.311191469123422e-10, 6.362359589419096e-10,
	6.414249531371391e-10, 6.466885438281477e-10, 6.52029265242335e-10,
	6.574497796711595e-10, 6.629528863437449e-10, 6.685415310821349e-10,
	6.742188168224279e-10, 6.799880150968072e-10, 6.858525785838827e-10,
	6.91816154849038e-10, 6.97882601412975e-10, 7.040560023057454e-10,
	7.103406862857415e-10, 7.167412469289477e-10, 7.23262564823922e-10,
	7.299098321433274e-10, 7.36688579904375e-10, 7.436047082795391e-10,
	7.506645203768892e-10, 7.578747599782539e-10, 7.652426538055458e-10,
	7.727759589838675e-10, 7.804830164881679e-10, 7.883728115028472e-10,
	7.964550417966955e-10, 8.047401954263357e-10, 8.132396393395169e-10,
	8.219657207674682e-10, 8.309318836890947e-10, 8.401528031399729e-10,
	8.496445407534143e-10, 8.594247256958435e-10, 8.695127661432599e-10,
	8.799300977056073e-10, 8.907004768313693e-10, 9.0185032933939e-10,
	9.134091670009051e-10, 9.254100887742333e-10, 9.378903882223966e-10,
	9.508922953177937e-10, 9.644638899862886e-10, 9.786602374481e-10,
	9.935448133101142e-10, 1.0091913119697182e-09, 1.0256859691519228e-09,
	1.0431305846498399e-09, 1.0616465149697269e-09, 1.081380035127533e-09,
	1.1025096747562618e-09, 1.1252564706432428e-09, 1.1498986477733707e-09,
	1.1767932423346918e-09, 1.2064090187897673e-09, 1.2393785886825987e-09,
	1.276584953890662e-09, 1.3193139264951536e-09, 1.3695434471115933e-09,
	1.4305498138471676e-09, 1.5083650345524237e-09, 1.6160853275510512e-09,
	1.792124814850057e-09,
}
var fe = [256]float64{
	1, 0.9381436808621765, 0.9004699299257477, 0.8717043323812047,
	0.8477855006239905, 0.8269932966430511, 0.808421651523009,
	0.7915276369724963, 0.7759568520401162, 0.7614633888498968,
	0.7478686219851957, 0.735038092431424, 0.7228676595935725,
	0.7112747608050765, 0.7001926550827886, 0.6895664961170784,
	0.6793505722647658, 0.6695063167319252, 0.6600008410790001,
	0.6508058334145714, 0.6418967164272664, 0.6332519942143664,
	0.6248527387036662, 0.6166821809152079, 0.6087253820796223,
	0.6009689663652326, 0.5934009016917338, 0.5860103184772684,
	0.5787873586028454, 0.5717230486648262, 0.5648091929124006,
	0.5580382822625879, 0.5514034165406417, 0.5448982376724401,
	0.5385168720028622, 0.5322538802630437, 0.5261042139836201,
	0.5200631773682339, 0.5141263938147489, 0.5082897764106432,
	0.5025495018413481, 0.4969019872415499, 0.49134386959403287,
	0.48587198734188525, 0.48048336393045454, 0.4751751930373777,
	0.4699448252839603, 0.4647897562504265, 0.459707615642138,
	0.45469615747461584, 0.44975325116275533, 0.44487687341454885,
	0.4400651008423542, 0.4353161032156369, 0.43062813728845917,
	0.4259995411430347, 0.4214287289976169, 0.41691418643300326,
	0.4124544659971615, 0.40804818315203273, 0.4036940125305306,
	0.3993906844752314, 0.39513698183329043, 0.3909317369847974,
	0.38677382908413793, 0.38266218149601006, 0.37859575940958107,
	0.37457356761590244, 0.3705946484351463, 0.36665807978151443,
	0.3627629733548181, 0.35890847294875006, 0.35509375286678774,
	0.3513180164374836, 0.34758049462163726, 0.3438804447045027,
	0.34021714906678024, 0.3365899140286778, 0.3329980687618092,
	0.32944096426413655, 0.32591797239355635, 0.3224284849560893,
	0.31897191284495735, 0.315547685227129, 0.3121552487741797,
	0.30879406693456024, 0.3054636192445903, 0.3021634006756935,
	0.2988929210155818, 0.29565170428126125, 0.29243928816189263,
	0.28925522348967775, 0.2860990737370769, 0.2829704145387808,
	0.2798688332369729, 0.2767939284485174, 0.27374530965280297,
	0.27072259679906, 0.26772541993204485, 0.26475341883506226,
	0.26180624268936303, 0.2588835497490163, 0.25598500703041543,
	0.2531102900156295, 0.25025908236886235, 0.24743107566532765,
	0.24462596913189213, 0.24184346939887724, 0.23908329026244915,
	0.23634515245705962, 0.23362878343743335, 0.23093391716962744,
	0.22826029393071676, 0.22560766011668415, 0.22297576805812028,
	0.22036437584335958, 0.21777324714870058, 0.2152021510753787,
	0.21265086199297834, 0.2101191593889883, 0.20760682772422212,
	0.20511365629383782, 0.2026394390937091, 0.20018397469191135,
	0.19774706610509893, 0.19532852067956327, 0.19292814997677138,
	0.19054576966319545, 0.18818119940425432, 0.18583426276219714,
	0.1835047870977675, 0.18119260347549634, 0.17889754657247836,
	0.17661945459049494, 0.17435816917135352, 0.1721135353153201,
	0.16988540130252766, 0.1676736186172502, 0.165478041874936,
	0.16329852875190182, 0.16113493991759203, 0.1589871389693142,
	0.15685499236936523, 0.15473836938446808, 0.15263714202744288,
	0.15055118500103992, 0.14848037564386682, 0.14642459387834497,
	0.1443837221606348, 0.14235764543247223, 0.1403462510748625,
	0.1383494288635803, 0.13636707092642894, 0.1343990717022137,
	0.13244532790138763, 0.13050573846833088, 0.1285802045452283,
	0.12666862943751078, 0.12477091858083104, 0.12288697950954522,
	0.1210167218266749, 0.11916005717532775, 0.11731689921155564,
	0.1154871635786336, 0.11367076788274438, 0.11186763167005638,
	0.11007767640518545, 0.10830082545103385, 0.10653700405000172,
	0.10478613930657024, 0.1030481601712578, 0.10132299742595369,
	0.09961058367063715, 0.09791085331149221, 0.09622374255043283,
	0.09454918937605587, 0.09288713355604357, 0.09123751663104017,
	0.08960028191003284, 0.08797537446727019, 0.08636274114075689,
	0.0847623305323681, 0.08317409300963235, 0.08159798070923742,
	0.0800339475423199, 0.07848194920160644, 0.07694194317048052,
	0.0754138887340584, 0.07389774699236475, 0.07239348087570872,
	0.07090105516237181, 0.06942043649872875, 0.06795159342193662,
	0.06649449638533979, 0.06504911778675376, 0.06361543199980735,
	0.06219341540854101, 0.06078304644547963, 0.05938430563342025,
	0.05799717563120064, 0.05662164128374284, 0.05525768967669701,
	0.05390531019604605, 0.052564494593071664, 0.051235237055126254,
	0.04991753428270636, 0.04861138557337948, 0.04731679291318155,
	0.04603376107617516, 0.04476229773294327, 0.043502413568888176,
	0.04225412241331624, 0.04101744138041482, 0.03979239102337412,
	0.03857899550307485, 0.03737728277295936, 0.03618728478193143,
	0.03500903769739742, 0.033842582150874344, 0.03268796350895954,
	0.03154523217289361, 0.030414443910466608, 0.029295660224637397,
	0.028188948763978632, 0.0270943837809558, 0.02601204664513422,
	0.024942026419731787, 0.023884420511558174, 0.02283933540638524,
	0.021806887504283584, 0.020787204072578117, 0.01978042433800974,
	0.018786700744696024, 0.017806200410911355, 0.01683910682603994,
	0.015885621839973156, 0.014945968011691148, 0.014020391403181943,
	0.013109164931254991, 0.012212592426255378, 0.0113310135978346,
	0.010464810181029982, 0.009614413642502213, 0.008780314985808977,
	0.007963077438017043, 0.007163353183634991, 0.006381905937319183,
	0.005619642207205489, 0.0048776559835424, 0.0041572951208338005,
	0.003460264777836907, 0.0027887987935740783, 0.002145967743718907,
	0.0015362997803015728, 0.0009672692823271743, 0.0004541343538414966,
}
