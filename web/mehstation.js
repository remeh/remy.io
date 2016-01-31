var player;
var player_target;

function onYouTubeIframeAPIReady() {
  player = new YT.Player('player', {
    height: '290',
    width: '100%',
    videoId: 'FxudzfhMTlU',
    events: {
      'onReady': onPlayerReady,
      //'onStateChange': onPlayerStateChange
    }
  });
}

function onPlayerReady(event) {
  player_target = event.target;
}

function onStart() {
  var o = document.querySelectorAll('#overplayer');
  var p = document.querySelectorAll('#player');
  console.log(o);
  console.log(p);
  if (o.length > 0 && p.length > 0 && player_target) {
    o[0].style.visibility = 'hidden';
    o[0].style.display = 'none';
    p[0].style.visibility = 'visible';
    p[0].style.display = 'inherit';
    p[0].style.height = '620px';
    player_target.playVideo();
    mixpanel.track('mehstation video start');
  }
}
